package main

import (
	"model"
	"time"

	"github.com/jinzhu/gorm"
)

type Cleaner struct {
	db             *gorm.DB
	config         *model.Config
	queryTimestamp int64
}

func CreateCleaner(cfg *model.Config) *Cleaner {
	db, err := gorm.Open(cfg.Driver, cfg.ConnectionString)

	if err != nil {
		panic(err)
	}

	service := new(Cleaner)
	service.db = db
	service.config = cfg

	return service
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
		service.db.Where(&model.Articles{FeedId: feed.Id}).Count(&articlesCount)

		if articlesCount > service.config.ArticlesMaxCount {
			go service.deleteArticles(feed.Id)
		}
	}
}

func (service *Cleaner) deleteArticles(feedId int) {
	// fixme
	service.db.
		Where("Date < ? AND IsBookmark=0 AND IsRead=1 AND FeedId=?", service.queryTimestamp, feedId).
		Delete(model.Articles{})
}
