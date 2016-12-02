package models

type Feeds struct {
	Id       uint   `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name     string `gorm:"column:Name"`
	Url      string `gorm:"column:Url"`
	UserId   uint   `gorm:"column:UserId;index"`
	Articles []Articles
}

type Articles struct {
	Id         uint   `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	FeedId     uint   `gorm:"column:FeedId;index"`
	Title      string `gorm:"column:Title"`
	Body       string `gorm:"column:Body;size:8192"`
	Link       string `gorm:"column:Link"`
	Date       int64  `gorm:"column:Date"`
	IsRead     bool   `gorm:"column:IsRead"`
	IsBookmark bool   `gorm:"column:IsBookmark"`
}

type Users struct {
	Id       uint   `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name     string `gorm:"column:Name"`
	Password string `gorm:"column:Password"`
	Settings Settings
	Feeds    []Feeds
}

type Settings struct {
	Id         uint `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	UserId     uint `gorm:"column:UserId;index"`
	UnreadOnly bool `gorm:"column:UnreadOnly"`
}
