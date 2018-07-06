package model

// Rss - structure for DB
type Feeds struct {
	Id       int        `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name     string     `gorm:"column:Name"`
	Url      string     `gorm:"column:Url"`
	UserId   int        `gorm:"column:UserId"`
	Articles []Articles `gorm:"ForeignKey:FeedId"`
}

func (Feeds) TableName() string {
	return "feeds"
}

type Articles struct {
	Id         int    `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	FeedId     int    `gorm:"column:FeedId;index"`
	Title      string `gorm:"column:Title"`
	Body       string `gorm:"column:Body;size:8192"`
	Link       string `gorm:"column:Link"`
	Date       int64  `gorm:"column:Date"`
	IsRead     bool   `gorm:"column:IsRead"`
	IsBookmark bool   `gorm:"column:IsBookmark"`
	//Feed       Feeds
}

func (Articles) TableName() string {
	return "articles"
}
