package models

// Rss - structure for DB
type Feeds struct {
	Id       uint       `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name     string     `gorm:"column:Name"`
	Url      string     `gorm:"column:Url"`
	UserId   uint       `gorm:"column:UserId"`
	Articles []Articles `gorm:"ForeignKey:FeedId"`
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
	//Feed       Feeds
}

func (Articles) TableName() string {
	return "articles"
}

type Users struct {
	Id                uint   `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name              string `gorm:"column:Name"`
	Password          string `gorm:"column:Password"`
	VkLogin           string `gorm:"column:VkLogin"`
	VkPassword        string `gorm:"column:VkPassword"`
	TwitterScreenName string `gorm:"column:TwitterScreenName"`
	VkNewsEnabled     bool   `gorm:"column:VkNewsEnabled"`
	Settings          Settings
	Feeds             []Feeds
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
	TwitterEnabled    bool `gorm:"column:TwitterEnabled"`
	TwitterSimpleVersion    bool `gorm:"column:TwitterSimpleVersion"`
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
	Id        int    `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	UserId    int    `gorm:"column:UserId;index"`
	GroupId   int    `gorm:"column:GroupId;index"`
	PostId    int    `gorm:"column:PostId;index"`
	Text      string `gorm:"column:Text"`
	Image     string `gorm:"column:Image"`
	Link      string `gorm:"column:Link"`
	Timestamp int64  `gorm:"column:Timestamp"`
}

func (VkNews) TableName() string {
	return "vknews"
}

type VkGroup struct {
	Id         int    `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Gid        int    `gorm:"column:Gid;index"`
	UserId     int    `gorm:"column:UserId;index"`
	Name       string `gorm:"column:Name"`
	LinkedName string `gorm:"column:LinkedName"`
	Image      string `gorm:"column:Image"`
}

func (VkGroup) TableName() string {
	return "vkgroups"
}

/* Twitter Models
============================================================================= */
type TwitterNews struct {
	Id          uint64 `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	UserId      int    `gorm:"column:UserId;index"`
	SourceId    int    `gorm:"column:SourceId;index"`
	Text        string `gorm:"column:Text"`
	ExpandedUrl string `gorm:"column:ExpandedUrl"`
	Image       string `gorm:"column:Image"`
}

func (TwitterNews) TableName() string {
	return "twitternews"
}

type TwitterSource struct {
	Id         int    `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	UserId     int    `gorm:"column:UserId;index"`
	Name       string `gorm:"column:Name"`
	ScreenName string `gorm:"column:ScreenName"`
	Url        string `gorm:"column:Url"`
	Image      string `gorm:"column:Image"`
}

func (TwitterSource) TableName() string {
	return "twittersource"
}
