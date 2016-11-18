package models

import "github.com/jinzhu/gorm"

// Rss - structure for DB
type Rss struct {
	ID       uint   `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	RssName  string `gorm:"column:Name"`
	RssURL   string `gorm:"column:Url"`
	Articles []RssArticle
}

func (Rss) TableName() string {
	return "feeds"
}

// RssArticle - article in feed
type RssArticle struct {
	ID         uint   `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	RssID      uint   `gorm:"column:FeedId;index"`
	Title      string `gorm:"column:Title"`
	Body       string `gorm:"column:Body"`
	Link       string `gorm:"column:Link"`
	Date       int64  `gorm:"column:Date"`
	IsRead     bool   `gorm:"column:IsRead"`
	IsBookmark bool   `gorm:"column:IsBookmark"`
}

func (RssArticle) TableName() string {
	return "articles"
}

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Password string
	Settings Settings
}

func (User) TableName() string {
	return "users"
}

type Settings struct {
	gorm.Model
	UserID     uint
	UnreadOnly bool
}

func (Settings) TableName() string {
	return "settings"
}
