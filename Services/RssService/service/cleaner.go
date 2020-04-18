package service

import (
	"newshub-rss-service/model"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Cleaner struct {
	db             *gorm.DB
	config         model.Config
	queryTimestamp int64
}

func CreateCleaner(cfg model.Config) *Cleaner {
	// fixme: держать соединение постоянно ненужно
	db, err := gorm.Open(cfg.Driver, cfg.ConnectionString)

	if err != nil {
		panic(err)
	}

	service := new(Cleaner)
	service.db = db
	service.config = cfg

	return service
}

func (service *Cleaner) Close() {
	if service.db != nil {
		service.db.Close()
	}
}

// CleanOldArticles - remove articles where create date less month
func (service Cleaner) Clean() {
	now := time.Now().Unix()
	month := int64(60 * 60 * 24 * 30)
	service.queryTimestamp = now - month

	service.preHandle()
}

func (service *Cleaner) preHandle() {
	var feeds []model.Feeds
	service.db.Find(&feeds)

	for _, feed := range feeds {
		var articlesCount int
		service.db.Model(&model.Articles{}).Where(&model.Articles{FeedId: feed.Id}).Count(&articlesCount)

		if articlesCount > service.config.ArticlesMaxCount {
			go service.deleteArticles(feed.Id)
		}
	}
}

func (service *Cleaner) deleteArticles(feedId int64) {
	// fixme
	service.db.
		Where("Date < ? AND IsBookmark=0 AND IsRead=1 AND FeedId=?", service.queryTimestamp, feedId).
		Delete(model.Articles{})
}
