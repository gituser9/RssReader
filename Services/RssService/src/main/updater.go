package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"encoding/xml"
	"io"
	"log"
	"model"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
)

const saveChanBufferSize = 10
const getChanBufferSize = 10

// Updater - service
type Updater struct {
	db       *gorm.DB
	config   *model.Config
	getChan  chan getData
	saveChan chan model.Articles
}

type getData struct {
	Rss  model.Feeds
	Body io.ReadCloser
}

// CreateUpdater - create and configure Updater struct
func CreateUpdater(cfg *model.Config) *Updater {
	service := new(Updater)
	service.config = cfg
	service.getChan = make(chan getData, getChanBufferSize)
	service.saveChan = make(chan model.Articles, saveChanBufferSize)

	go service.updateFeedRunner()
	go service.saveArticle()

	return service
}

// Update - get new feeds for users
func (service *Updater) Update() {
	db, err := gorm.Open(service.config.Driver, service.config.ConnectionString)

	if err != nil {
		log.Println("open db error:", err.Error())
		return
	}

	defer db.Close()
	service.db = db

	settings := make([]model.Settings, 0)
	service.db.Where(&model.Settings{RssEnabled: true}).Find(&settings)

	if settings == nil || len(settings) == 0 {
		return
	}
	for _, item := range settings {
		var feeds []model.Feeds
		service.db.Where(&model.Feeds{UserId: item.UserId}).Find(&feeds)

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
		log.Println(err.Error())
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
			err := decoder.Decode(&xmlModel)

			if err != nil {
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
	links := make(map[string]bool, len(rss.Articles))

	for _, article := range rss.Articles {
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
