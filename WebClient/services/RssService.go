package services

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"newshub-server/models"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/net/html/charset"
)

// RssService - service
type RssService struct {
	db         *gorm.DB
	config     *models.Config
	UnreadOnly bool
}

func NewRssService(config *models.Config) *RssService {
	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println(err)
	}

	return &RssService{db: db, config: config}
}

func (service *RssService) SetDb(db *gorm.DB) {
	service.db = db
}

func (service *RssService) SetConfig(cfg *models.Config) {
	service.config = cfg
}

// GetRss - get all rss
func (service *RssService) GetRss(id int64) []models.Feed {
	var rss []models.Feeds
	service.dbp().
		Preload("Articles", "IsRead=?", "0").
		Where(&models.Feeds{UserId: id}).
		Find(&rss)
	feeds := make([]models.Feed, len(rss))
	var wg sync.WaitGroup

	for i, item := range rss {
		wg.Add(1)
		go func(item models.Feeds, i int) {
			count := len(item.Articles)
			item.Articles = nil
			feeds[i] = models.Feed{Feed: item, ArticlesCount: count, ExistUnread: count > 0}

			wg.Done()
		}(item, i)
	}

	wg.Wait()

	return feeds
}

// GetArticles - get articles for rss by id
func (service *RssService) GetArticles(id int64, userId int64, page int) *models.ArticlesJSON {
	var articles []models.Articles
	var count int
	offset := service.config.PageSize * (page - 1)
	whereObject := models.Articles{FeedId: id}

	query := service.dbp().Where(&whereObject).
		Select("Id, Title, IsBookmark, IsRead, Link").
		Limit(service.config.PageSize).
		Offset(offset).
		Order("Id desc")
	queryCount := service.dbp().Model(&whereObject).Where(&whereObject)

	var settings models.Settings
	service.dbp().Where(models.Settings{UserId: userId}).Find(&settings)

	if settings.UnreadOnly {
		whereNotObject := models.Articles{IsRead: true}
		query = query.Not(&whereNotObject)
		queryCount = queryCount.Not(&whereNotObject)
	}

	query.Find(&articles)
	queryCount.Count(&count)

	return &models.ArticlesJSON{Articles: articles, Count: count}
}

// GetArticle - get one article
func (service *RssService) GetArticle(id int64, feedId int64, userId int64) *models.Articles {
	rss := service.GetRss(userId)
	log.Println(rss)

	if len(rss) == 0 {
		return nil
	}

	// get article
	var article models.Articles
	service.dbp().Where(&models.Articles{Id: id, FeedId: feedId}).First(&article)

	var settings models.Settings // todo: to func
	service.dbp().Where(models.Settings{UserId: userId}).Find(&settings)

	// update state
	article.IsRead = true
	service.dbp().Save(&article)

	if settings.MarkSameRead {
		go service.markSameArticles(article.Link, article.FeedId)
	}

	return &article
}

// Import - import OPML file
func (service *RssService) Import(data []byte, userId int64) {
	// parse opml
	var opml models.OPML
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&opml)

	if err != nil {
		log.Println("OPML import error: ", err.Error())
		return
	}

	dbExec(func(db *gorm.DB) {
		for _, outline := range opml.Outlines {
			feed := models.Feeds{
				Name:   outline.Title,
				Url:    outline.URL,
				UserId: userId,
			}
			db.Save(&feed)
		}
	})
}

// Export - export feeds to OPML file
func (service *RssService) Export(userId int64) {
	// get data from DB
	var rss []models.Feeds
	service.dbp().Where(&models.Feeds{UserId: userId}).Find(&rss)
	opml := models.OPML{
		HeadText: "Feeds",
		Version:  1.1,
	}

	// create array of structures
	for _, feed := range rss {
		outline := models.OPMLOutline{
			Title: feed.Name,
			URL:   feed.Url,
			Text:  feed.Name,
		}
		opml.Outlines = append(opml.Outlines, outline)
	}

	// create OPML file
	xmlString, _ := xml.Marshal(opml)
	opmlBytes, _ := ioutil.ReadFile(service.config.FilePath)
	var conf models.Config
	json.Unmarshal(opmlBytes, &conf)

	if len(conf.OPMLPath) > 0 {
		go ioutil.WriteFile(conf.OPMLPath+"/rss.opml", xmlString, 0777)
	}

	// todo: clean old
	fileName := uuid.New().String()
	ioutil.WriteFile(fmt.Sprintf("./static/%s.opml", fileName), xmlString, 0777)
}

// AddFeed - add new feed
func (service *RssService) AddFeed(url string, userId int64) {
	// get rss xml
	response, err := http.Get(url)

	if err != nil {
		log.Println("Get XML error: ", err.Error())
		return
	}

	defer response.Body.Close()

	// parse feed xml and create structure
	var xmlModel models.XMLFeed
	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&xmlModel)

	if err != nil {
		log.Println("XML unmarshall error on URL: ", url, err.Error())
		return
	}

	// insert in DB
	err = service.dbp().Create(&models.Feeds{Url: url, UserId: userId, Name: xmlModel.RssName}).Error
	// todo: send message for update

	if err != nil {
		log.Println("insert error", err.Error())
	}
}

// Delete - remove feed
func (service *RssService) Delete(id int64, userId int64) {
	feed := service.GetRss(userId)

	if len(feed) == 0 {
		return
	}

	service.dbp().Where(models.Articles{FeedId: id}).Delete(models.Articles{})
	service.dbp().Delete(models.Feeds{Id: id})
}

// SetNewName - update feed name
func (service *RssService) SetNewName(data models.FeedUpdateData, userId int64) {
	feed := models.Feeds{}
	service.dbp().Where(&models.Feeds{Id: data.FeedId, UserId: userId}).First(&feed)

	if feed.Id == 0 {
		return
	}

	if data.IsReadAll {
		service.dbp().Model(&models.Articles{}).
			Where(&models.Articles{FeedId: feed.Id}).
			Not(&models.Articles{IsRead: true}).
			UpdateColumn(models.Articles{IsRead: true})
	}
	if data.Name != "" {
		feed.Name = data.Name
		service.dbp().Save(&feed)
	}
}

// GetBookmarks - get all bookmarks
func (service *RssService) GetBookmarks(page int, userId int64) *models.ArticlesJSON {
	var articles []models.Articles
	whereCond := "articles.IsBookmark = true and feeds.UserId = ?"
	offset := service.config.PageSize * (page - 1)
	var count int

	service.dbp().Where(whereCond, userId).
		Joins("join feeds on articles.FeedId = feeds.Id").
		Select("Id, Title, IsBookmark, IsRead").
		Limit(service.config.PageSize).
		Offset(offset).
		Order("Id desc").
		Find(&articles)
	service.dbp().Model(&models.Articles{}).Where(whereCond, userId).
		Joins("join feeds on articles.FeedId = feeds.Id").Count(&count)

	return &models.ArticlesJSON{Articles: articles, Count: count}
}

// Search - search articles by title or body
func (service *RssService) Search(searchString string, isBookmark bool, feedId int64, userId int64) *models.ArticlesJSON {
	var articles []models.Articles
	query := service.dbp().
		Joins("join feeds on articles.FeedId = feeds.Id").
		Select("Id, Title, IsBookmark, IsRead, Link").
		Where("(articles.Title LIKE ? OR articles.Body LIKE ?) and feeds.UserId = ?", "%"+searchString+"%", "%"+searchString+"%", userId)

	if feedId != 0 {
		query = query.Where(&models.Articles{Id: feedId})
	}
	if isBookmark {
		query = query.Where("IsBookmark = 1")
	}

	query.Find(&articles)

	return &models.ArticlesJSON{Articles: articles}
}

func (service *RssService) ArticleUpdate(userId int64, data models.ArticlesUpdateData) {
	whereCond := "articles.Id = ? and feeds.UserId = ?"
	article := models.Articles{}
	service.dbp().
		Joins("join feeds on articles.FeedId = feeds.Id").
		Where(whereCond, data.ArticleId, userId).
		First(&article)
	if article.Id != 0 {
		article.IsBookmark = data.IsBookmark
		article.IsRead = data.IsRead
		service.dbp().Save(&article)
	}
}

func (service *RssService) markSameArticles(url string, feedID int64) {
	updateModel := models.Articles{IsRead: true}
	service.dbp().Model(&models.Articles{}).Where(&models.Articles{Link: url}).
		Not(&models.Articles{Id: feedID}).
		UpdateColumn(&updateModel)
}

func (service *RssService) dbp() *gorm.DB {
	if service.db == nil {
		var err error
		service.db, err = gorm.Open(service.config.Driver, service.config.ConnectionString)

		if err != nil {
			log.Println("DB open error: ", err)
			os.Exit(1)
		}
	}

	return service.db
}
