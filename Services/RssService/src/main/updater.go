package main


import (
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"

	"model"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"log"
	"io"
	"net/http"
	"sync"
	"time"
)

// Updater - service
type Updater struct {
	db          *gorm.DB
	config      *model.Config
	saveChan chan saveData
}

type saveData struct {
	Rss model.Feeds
	Body io.ReadCloser
}

func CreateUpdater(cfg *model.Config) *Updater {
	db, err := gorm.Open(cfg.Driver, cfg.ConnectionString)

	if err != nil {
		panic(err)
	}

	service := new(Updater)
	service.db = db
	service.config = cfg
	service.saveChan = make(chan saveData, 0)

	go service.updateFeedRunner()

	return service
}

func (service *Updater) Update() {
	settings := make([]model.Settings, 0)
	service.db.Where(&model.Settings{RssEnabled: true}).Find(&settings)

	for _, item := range settings {
		var feeds []model.Feeds
		service.db.Where(&model.Feeds{UserId: item.UserId}).Find(&feeds)

		for _, feed := range feeds {
			rssBody, err := service.getFeedBody(feed.Url)

			if err != nil {
				log.Println("get rss error: ", err.Error())
				continue
			}
			if rssBody == nil {
				continue
			}

			service.saveChan <- saveData{Body: rssBody, Rss: feed}
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

func (service Updater) updateFeedRunner() {
	for {
		select {
		case data := <-service.saveChan:
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
	links := make([]string, len(rss.Articles))
	var wg sync.WaitGroup

	for i, article := range rss.Articles {
		links[i] = article.Link
	}

	newArticles := make([]model.Articles, len(xmlModel.Articles))

	for i, article := range xmlModel.Articles {
		wg.Add(1)

		go func(i int, article model.XMLArticle) {
			defer wg.Done()
			isExist := false

			// todo: go and channels
			for _, item := range links {
				if item == article.Link {
					isExist = true
					break
				}
			}
			if !isExist {
				newArticle := service.rssArticleFromXML(&article)
				newArticle.FeedId = rss.Id
				newArticles[i] = newArticle
			}

		}(i, article)
	}

	wg.Wait()

	for _, article := range newArticles {
		if article.FeedId != 0 {
			service.db.Create(&article)
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