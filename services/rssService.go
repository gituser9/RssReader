package services

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

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

// GetRss - get all rss
func (service *RssService) GetRss(id uint) []models.Feed {
	var rss []models.Rss
	service.dbp().Where(&models.Rss{UserId: id}).Preload("Articles", "IsRead=?", "0").Find(&rss)
	feeds := make([]models.Feed, len(rss))
	var wg sync.WaitGroup

	for i, item := range rss {
		wg.Add(1)
		go func(item models.Rss, i int) {
			count := len(item.Articles)
			feeds[i] = models.Feed{Rss: item, ArticlesCount: count, ExistUnread: count > 0}

			wg.Done()
		}(item, i)
	}

	wg.Wait()

	return feeds
}

// GetArticles - get articles for rss by id
func (service *RssService) GetArticles(id uint, page int) *models.ArticlesJSON {
	var articles []models.RssArticle
	var count int
	offset := service.config.PageSize * (page - 1)
	whereObject := models.RssArticle{RssID: id}

	query := service.dbp().Where(&whereObject).
		Select("Id, Title, IsBookmark, IsRead").
		Limit(service.config.PageSize).
		Offset(offset).
		Order("Id desc")
	queryCount := service.dbp().Model(&whereObject).Where(&whereObject)

	if service.AppSettings.UnreadOnly {
		whereNotObject := models.RssArticle{IsRead: true}
		query = query.Not(&whereNotObject)
		queryCount = queryCount.Not(&whereNotObject)
	}

	query.Find(&articles)
	queryCount.Count(&count)

	return &models.ArticlesJSON{Articles: articles, Count: count}
}

// GetArticle - get one article
func (service *RssService) GetArticle(id uint) *models.RssArticle {
	// get article
	var article models.RssArticle
	service.dbp().First(&article, id)

	// update state
	article.IsRead = true
	service.dbp().Save(&article)

	if service.AppSettings.MarkSameRead {
		go service.markSameArticles(article.Link, article.RssID)
	}

	return &article
}

// UpdateAllFeeds - update all feeds
func (service *RssService) UpdateAllFeeds() {
	var feeds []models.Rss
	service.dbp().Find(&feeds)
	var wg sync.WaitGroup

	for _, feed := range feeds {
		wg.Add(1)
		go service.UpdateFeed(feed.RssURL, &wg, feed.UserId)
	}

	wg.Wait()
}

// UpdateFeed - update one feed
func (service *RssService) UpdateFeed(url string, wg *sync.WaitGroup, userId uint) {
	// get xml by url
	defer wg.Done()
	rssBody, err := service.getFeedBody(url)

	if err != nil {
		log.Println("get rss error: ", err.Error())
		return
	}

	// get feed from DB by url, if not - add
	defer rssBody.Close()
	var rss models.Rss
	service.dbp().Preload("Articles").Where(&models.Rss{RssURL: url, UserId: userId}).First(&rss)

	if rss.RssURL == "" {
		service.AddFeed(url, userId)
		return
	}

	// unmarshal xml
	var xmlModel models.XMLFeed
	decoder := xml.NewDecoder(rssBody)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&xmlModel)

	if err != nil {
		log.Println("unmarshal error: " + err.Error())
		return
	}

	// update DB
	service.updateArticles(rss, xmlModel)
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

	var wg sync.WaitGroup

	// update feeds
	for _, item := range opml.Outlines {
		wg.Add(1)
		go service.UpdateFeed(item.URL, &wg, userId)
	}

	wg.Wait()
}

// Export - export feeds to OPML file
func (service *RssService) Export(userId uint) {
	// get data from DB
	var rss []models.Rss
	service.dbp().Where(&models.Rss{UserId: userId}).Find(&rss)
	opml := models.OPML{
		HeadText: "Feeds",
		Version:  1.1,
	}

	// create array of structures
	for _, feed := range rss {
		outline := models.OPMLOutline{
			Title: feed.RssName,
			URL:   feed.RssURL,
			Text:  feed.RssName,
		}
		opml.Outlines = append(opml.Outlines, outline)
	}

	// create OPML file
	xmlString, _ := xml.Marshal(opml)
	bytes, _ := ioutil.ReadFile(service.config.FilePath)
	var conf models.Config
	json.Unmarshal(bytes, &conf)

	if len(conf.OPMLPath) > 0 {
		go ioutil.WriteFile(conf.OPMLPath+"/rss.opml", xmlString, 0777)
	}

	ioutil.WriteFile("./static/rss.opml", xmlString, 0777)
}

// AddFeed - add new feed
func (service *RssService) AddFeed(url string, userId uint) {
	// get rss xml
	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		log.Println(err.Error())
		return
	}

	if err != nil {
		log.Println("Get XML error: ", err.Error())
		return
	}

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
	dbModel := service.fromXMLToDbStructure(&xmlModel)
	dbModel.RssURL = url
	dbModel.UserId = userId
	service.dbp().Create(&dbModel)

	if err != nil {
		log.Println("insert error", err.Error())
	}
}

// Delete - remove feed
func (service *RssService) Delete(id uint) {
	service.dbp().Where(models.RssArticle{RssID: id}).Delete(models.RssArticle{})
	service.dbp().Delete(models.Rss{ID: id})
}

// SetNewName - update feed name
func (service *RssService) SetNewName(newName string, id uint) {
	var feed models.Rss
	service.dbp().Where(models.Rss{ID: id}).First(&feed)
	feed.RssName = newName
	service.dbp().Save(&feed)
}

// ToggleBookmark - toggle article status as bookmark or not bookmark
func (service *RssService) ToggleBookmark(id uint, isBookmark bool) {
	// fixme
	var article models.RssArticle
	service.dbp().Where(&models.RssArticle{ID: id}).Find(&article)
	updateArticles := service.dbp().Model(&models.RssArticle{ID: article.ID}) // fixme: (should be link)

	if isBookmark {
		updateArticles.UpdateColumn(&models.RssArticle{IsBookmark: isBookmark})
	} else {
		updateArticles.UpdateColumn("IsBookmark", "0")
	}
}

// GetBookmarks - get all bookmarks
func (service *RssService) GetBookmarks(page int64) *models.ArticlesJSON {
	articles := []models.RssArticle{}
	whereObject := models.RssArticle{IsBookmark: true}
	offset := service.config.PageSize * (int(page) - 1)
	var count int

	service.dbp().Where(&whereObject).Find(&articles).
		Limit(service.config.PageSize).
		Offset(offset).
		Order("Id desc")
	service.dbp().Model(&models.RssArticle{}).Where(&whereObject).Count(&count)

	return &models.ArticlesJSON{Articles: articles, Count: count}
}

// MarkReadAll - mark read all articles by feed id
func (service *RssService) MarkReadAll(id uint) {
	service.dbp().Model(&models.RssArticle{}).
		Where(&models.RssArticle{RssID: id}).
		Not(&models.RssArticle{IsRead: true}).
		UpdateColumn(models.RssArticle{IsRead: true})
}

// CleanOldArticles - remove articles where create date less mounth
func (service *RssService) CleanOldArticles() {
	now := time.Now().Unix()
	month := int64(60 * 60 * 24 * 30)
	queryTimestamp := now - month
	// fixme
	service.dbp().Where("Date < ? AND IsBookmark=0 AND IsRead=1", queryTimestamp).Delete(models.RssArticle{})
}

// Search - search articles by title or body
func (service *RssService) Search(searchString string, isBookmark bool, feedID uint) *models.ArticlesJSON {
	// fixme
	bm := 0

	if isBookmark {
		bm = 1
	}

	var articles []models.RssArticle
	query := service.dbp().
		Select("Id, Title, IsBookmark, IsRead").
		Where("IsBookmark = ? AND (Title LIKE ? OR Body LIKE ?)", bm, "%"+searchString+"%", "%"+searchString+"%")

	if feedID != 0 {
		query = query.Where(&models.RssArticle{RssID: feedID})
	}

	query.Find(&articles)

	return &models.ArticlesJSON{Articles: articles}
}

// ToggleAsRead - set read or unread status for article
func (service *RssService) ToggleAsRead(id uint, isRead bool) {
	// fixme: begin
	updateArticle := service.dbp().Model(&models.RssArticle{ID: id})

	if isRead {
		updateArticle.UpdateColumn(models.RssArticle{IsRead: true})
	} else {
		updateArticle.UpdateColumn("IsRead", "0")
	}
	// fixme: end
}

// Backup - backup DB
func (service *RssService) Backup() {
	if len(service.config.DbBackupPath) == 0 {
		return
	}
}

/*==============================================================================
	Private
==============================================================================*/
// fromXMLToDbStructure - create Rss structure from XMLFeed structure
func (service *RssService) fromXMLToDbStructure(xmlModel *models.XMLFeed) *models.Rss {
	feed := models.Rss{
		RssName:  xmlModel.RssName,
		Articles: make([]models.RssArticle, 0),
	}

	for _, article := range xmlModel.Articles {
		rssArticle := service.rssArticleFromXML(&article)
		feed.Articles = append(feed.Articles, rssArticle)
	}

	return &feed
}

// rssArticleFromXML - create RssArticle from XMLArticle
func (service *RssService) rssArticleFromXML(xmlArticle *models.XMLArticle) models.RssArticle {
	rssArticle := models.RssArticle{
		Body:   xmlArticle.Description,
		Title:  xmlArticle.Title,
		Link:   xmlArticle.Link,
		Date:   time.Now().Unix(),
		IsRead: false,
	}

	return rssArticle
}

func (service *RssService) markSameArticles(url string, feedID uint) {
	updateModel := models.RssArticle{IsRead: true}
	service.dbp().Model(&models.RssArticle{}).Where(&models.RssArticle{Link: url}).
		Not(&models.RssArticle{RssID: feedID}).
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

func (service *RssService) startTimers(config *models.Config) {
	// todo: service must bo one only
	// todo: DB backup timer (24 hours)

	updateTime := time.Duration(service.AppSettings.UpdateMinutes) * time.Minute
	updateTimer := time.NewTicker(updateTime).C
	weekTimer := time.NewTicker(time.Hour * 168).C // week

	// on start
	go service.UpdateAllFeeds()

	for {
		select {
		case <-updateTimer:
			service.UpdateAllFeeds()
		case <-weekTimer:
			service.CleanOldArticles()
			//service.Export()
		}
	}
}

func (service *RssService) getFeedBody(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if response.StatusCode == 404 {
		log.Println(url, "404")
		return nil, err
	}

	return response.Body, nil
}

func (service *RssService) updateArticles(rss models.Rss, xmlModel models.XMLFeed) {
	links := make([]string, len(rss.Articles))
	var wg sync.WaitGroup

	for i, article := range rss.Articles {
		wg.Add(1)
		go func(i int, article models.RssArticle) {
			links[i] = article.Link
			wg.Done()
		}(i, article)
	}

	wg.Wait()
	newArticles := make([]models.RssArticle, len(xmlModel.Articles))

	for i, article := range xmlModel.Articles {
		wg.Add(1)
		go func(i int, article models.XMLArticle) {
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
				newArticle.RssID = rss.ID
				newArticles[i] = newArticle
			}

		}(i, article)
	}

	wg.Wait()

	for i, article := range newArticles {
		if article.RssID != 0 {
			service.dbp().Create(&newArticles[i])
		}
	}
}

func (service *RssService) reReadConfig() {
	bytes, err := ioutil.ReadFile(service.config.FilePath)

	if err != nil {
		log.Println("Read config file error")
	} else {
		json.Unmarshal(bytes, &service.config)
	}
}
