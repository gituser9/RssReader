package services

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/net/html/charset"

	"../models"
)

// RssService - service
type RssService struct {
	db          *gorm.DB
	config      *models.Config
	UnreadOnly  bool
	AppSettings models.AppSettings
}

// Init - create new struct pointer with collection
func (service *RssService) Init(config *models.Config) *RssService {
	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println(err)
	}

	// set default settings
	settings := models.AppSettings{
		MarkSameRead:  true,
		UpdateMinutes: config.UpdateMinutes,
	}

	return &RssService{db: db, config: config, AppSettings: settings}
}

func (service *RssService) SetDb(db *gorm.DB) {
	service.db = db
}

func (service *RssService) SetConfig(cfg *models.Config) {
	service.config = cfg
}

// GetRss - get all rss
func (service *RssService) GetRss(id uint) []models.Feed {
	var rss []models.Feeds
	service.dbp().Preload("Articles", "IsRead=?", "0").Where(&models.Feeds{UserId: id}).Find(&rss)
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
func (service *RssService) GetArticles(id uint, userId uint, page int) *models.ArticlesJSON {
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
func (service *RssService) GetArticle(id uint, userId uint) *models.Articles {
	// get article
	var article models.Articles
	service.dbp().First(&article, id)

	var settings models.Settings
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
func (service *RssService) Import(data []byte, userId uint) {
	// parse opml
	var opml models.OPML
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&opml)

	if err != nil {
		log.Println("OPML import error: ", err.Error())
		return
	}

	// todo: send message for update
}

// Export - export feeds to OPML file
func (service *RssService) Export(userId uint) {
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

	ioutil.WriteFile("./static/rss.opml", xmlString, 0777)
}

// AddFeed - add new feed
func (service *RssService) AddFeed(url string, userId uint) {
	// get rss xml
	response, err := http.Get(url)

	if err != nil {
		log.Println(err.Error())
		return
	}

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
	service.dbp().Create(&models.Feeds{Url: url, UserId: userId, Name: xmlModel.RssName})
	// todo: send message for update

	if err != nil {
		log.Println("insert error", err.Error())
	}
}

// Delete - remove feed
func (service *RssService) Delete(id uint) {
	service.dbp().Where(models.Articles{FeedId: id}).Delete(models.Articles{})
	service.dbp().Delete(models.Feeds{Id: id})
}

// SetNewName - update feed name
func (service *RssService) SetNewName(newName string, id uint) {
	var feed models.Feeds
	service.dbp().Where(models.Feeds{Id: id}).First(&feed)
	feed.Name = newName
	service.dbp().Save(&feed)
}

// ToggleBookmark - toggle article status as bookmark or not bookmark
func (service *RssService) ToggleBookmark(id uint, isBookmark bool) {
	// fixme
	var article models.Articles
	service.dbp().Where(&models.Articles{Id: id}).Find(&article)
	updateArticles := service.dbp().Model(&models.Articles{Id: article.Id}) // fixme: (should be link)

	if isBookmark {
		updateArticles.UpdateColumn(&models.Articles{IsBookmark: isBookmark})
	} else {
		updateArticles.UpdateColumn("IsBookmark", "0")
	}
}

// GetBookmarks - get all bookmarks
func (service *RssService) GetBookmarks(page int64, userId uint) *models.ArticlesJSON {
	var articles []models.Articles
	//whereObject := models.Articles{IsBookmark: true, Feed: models.Feeds{UserId: userId}}
	whereObject := models.Articles{IsBookmark: true}
	offset := service.config.PageSize * (int(page) - 1)
	var count int

	service.dbp().Where(&whereObject).
		Select("Id, Title, IsBookmark, IsRead").
		Limit(service.config.PageSize).
		Offset(offset).
		Order("Id desc").
		Find(&articles)
	service.dbp().Model(&models.Articles{}).Where(&whereObject).Count(&count)

	return &models.ArticlesJSON{Articles: articles, Count: count}
}

// MarkReadAll - mark read all articles by feed id
func (service *RssService) MarkReadAll(id uint) {
	service.dbp().Model(&models.Articles{}).
		Where(&models.Articles{FeedId: id}).
		Not(&models.Articles{IsRead: true}).
		UpdateColumn(models.Articles{IsRead: true})
}

// Search - search articles by title or body
func (service *RssService) Search(searchString string, isBookmark bool, feedId uint) *models.ArticlesJSON {
	var articles []models.Articles
	query := service.dbp().
		Select("Id, Title, IsBookmark, IsRead, Link").
		Where("(Title LIKE ? OR Body LIKE ?)", "%"+searchString+"%", "%"+searchString+"%")

	if feedId != 0 {
		query = query.Where(&models.Articles{Id: feedId})
	}
	if isBookmark {
		query = query.Where("IsBookmark = 1")
	}

	query.Find(&articles)

	return &models.ArticlesJSON{Articles: articles}
}

// ToggleAsRead - set read or unread status for article
func (service *RssService) ToggleAsRead(id uint, isRead bool) {
	// fixme: begin
	updateArticle := service.dbp().Model(&models.Articles{Id: id})

	if isRead {
		updateArticle.UpdateColumn(models.Articles{IsRead: true})
	} else {
		updateArticle.UpdateColumn("IsRead", "0")
	}
	// fixme: end
}

/*==============================================================================
	Private
==============================================================================*/

func (service *RssService) markSameArticles(url string, feedID uint) {
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
