package models

// Rss - structure for DB
type Feeds struct {
	Id       uint   `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name     string `gorm:"column:Name"`
	Url      string `gorm:"column:Url"`
	UserId   uint   `gorm:"column:UserId"`
	Articles []Articles
}

func (Feeds) TableName() string {
	return "feeds"
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
	Feed       Feeds
}

func (Articles) TableName() string {
	return "articles"
}

type Users struct {
	Id            uint   `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name          string `gorm:"column:Name"`
	Password      string `gorm:"column:Password"`
	VkLogin       string `gorm:"column:VkLogin"`
	VkPassword    string `gorm:"column:VkPassword"`
	VkNewsEnabled bool   `gorm:"column:VkNewsEnabled"`
	Settings      Settings
	Feeds         []Feeds
}

func (Users) TableName() string {
	return "users"
}

type Settings struct {
	Id                uint `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	UserId            uint `gorm:"column:UserId;index"`
	UnreadOnly        bool `gorm:"column:UnreadOnly"`
	MarkSameRead      bool `gorm:"column:MarkSameRead"`
	RssEnabled        bool `gorm:"column:RssEnabled"`
	VkNewsEnabled     bool `gorm:"column:VkNewsEnabled"`
	ShowPreviewButton bool `gorm:"column:ShowPreviewButton"`
	ShowTabButton     bool `gorm:"column:ShowTabButton"`
	ShowReadButton    bool `gorm:"column:ShowReadButton"`
}

func (Settings) TableName() string {
	return "settings"
}

/* Vk Models
============================================================================= */
type VkNews struct {
	Id    int
	Text  string
	Image string
}

func (VkNews) TableName() string {
	return "vknews"
}

type VkGroup struct {
	Id     int
	UserId uint
	Name   string
}

func (VkGroup) TableName() string {
	return "vkgroup"
}
