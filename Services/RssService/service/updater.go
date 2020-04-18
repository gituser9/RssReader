package service

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"newshub-rss-service/model"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/net/html/charset"
)

const saveChanBufferSize = 20
const getChanBufferSize = 20

// Updater - service
type Updater struct {
	db       *gorm.DB
	config   model.Config
	getChan  chan getData
	saveChan chan model.Articles
}

type getData struct {
	Rss  model.Feeds
	Body io.ReadCloser
}

// CreateUpdater - create and configure Updater struct
func CreateUpdater(cfg model.Config) *Updater {
	// fixme: держать соединение постоянно ненужно
	db, err := gorm.Open(cfg.Driver, cfg.ConnectionString)

	if err != nil {
		panic(err)
	}

	service := new(Updater)
	service.db = db
	service.config = cfg
	service.getChan = make(chan getData, getChanBufferSize)
	service.saveChan = make(chan model.Articles, saveChanBufferSize)

	go service.updateFeedRunner()
	go service.saveArticle()

	return service
}

func (service *Updater) Close() {
	if service.db != nil {
		service.db.Close()
	}
}

// Update - get new feeds for users
func (service *Updater) Update() {
	userIds, err := getUserIds(service.db)

	if err != nil {
		return
	}

	for _, id := range userIds {
		var feeds []model.Feeds

		// todo: preload articles
		if err := service.db.Where(&model.Feeds{UserId: id}).Find(&feeds).Error; err != nil {
			log.Printf("get feeds for user %d error: %s", id, err)
			continue
		}
		if feeds == nil || len(feeds) == 0 {
			continue
		}
		for _, feed := range feeds {
			rssBody, err := service.getFeedBody(feed.Url)

			if err != nil {
				log.Println("get rss error: ", err.Error())
				continue
			}
			if rssBody == nil {
				continue
			}

			service.getChan <- getData{Body: rssBody, Rss: feed}
		}
	}
}

func (service Updater) getFeedBody(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)

	if err != nil {
		log.Println("get feed error:", err.Error())
		return nil, err
	}
	if response.StatusCode == 404 {
		log.Println(url, "404")
		return nil, nil
	}

	return response.Body, nil
}

func (service *Updater) updateFeedRunner() {
	for {
		select {
		case data := <-service.getChan:
			var xmlModel model.XMLFeed
			decoder := xml.NewDecoder(data.Body)
			decoder.CharsetReader = charset.NewReaderLabel

			if err := decoder.Decode(&xmlModel); err != nil {
				data.Body.Close()
				continue
			}

			// update DB
			go service.updateArticles(data.Rss, xmlModel)
			data.Body.Close()
		}
	}
}

func (service *Updater) updateArticles(rss model.Feeds, xmlModel model.XMLFeed) {
	articles := []model.Articles{}
	service.db.
		Where(&model.Articles{FeedId: rss.Id}).
		Select("Link").
		Find(&articles)
	links := make(map[string]bool, len(articles))

	for _, article := range articles {
		links[article.Link] = true
	}
	for _, article := range xmlModel.Articles {
		if _, isExist := links[article.Link]; !isExist {
			newArticle := service.rssArticleFromXML(&article)
			newArticle.FeedId = rss.Id

			service.saveChan <- newArticle
		}
	}
}

// rssArticleFromXML - create RssArticle from XMLArticle
func (service *Updater) rssArticleFromXML(xmlArticle *model.XMLArticle) model.Articles {
	rssArticle := model.Articles{
		Body:   xmlArticle.Description,
		Title:  xmlArticle.Title,
		Link:   xmlArticle.Link,
		Date:   time.Now().Unix(),
		IsRead: false,
	}

	return rssArticle
}

func (service *Updater) saveArticle() {
	for {
		select {
		case article := <-service.saveChan:
			if article.FeedId != 0 {
				service.db.Create(&article)
			}
		}
	}
}

func getUserIds(db *gorm.DB) ([]int64, error) {
	ids := make([]int64, 0)
	err := db.
		Model(&model.Settings{}).
		Where(&model.Settings{RssEnabled: true}).
		Pluck("UserId", &ids).
		Error

	if err != nil {
		return nil, err
	}

	return ids, nil
}
