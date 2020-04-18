package dao

type User struct {
	Id                int64    `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name              string   `gorm:"column:Name"`
	Password          string   `gorm:"column:Password" json:"-"`
	VkLogin           string   `gorm:"column:VkLogin"`
	VkPassword        string   `gorm:"column:VkPassword"`
	TwitterScreenName string   `gorm:"column:TwitterScreenName"`
	VkNewsEnabled     bool     `gorm:"column:VkNewsEnabled"`
	Settings          Settings `gorm:"ForeignKey:UserId"`
}

func (User) TableName() string {
	return "users"
}

type Settings struct {
	Id                   int64 `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	UserId               int64 `gorm:"column:UserId;index"`
	UnreadOnly           bool  `gorm:"column:UnreadOnly"`
	MarkSameRead         bool  `gorm:"column:MarkSameRead"`
	RssEnabled           bool  `gorm:"column:RssEnabled"`
	VkNewsEnabled        bool  `gorm:"column:VkNewsEnabled"`
	TwitterEnabled       bool  `gorm:"column:TwitterEnabled"`
	TwitterSimpleVersion bool  `gorm:"column:TwitterSimpleVersion"`
	ShowPreviewButton    bool  `gorm:"column:ShowPreviewButton"`
	ShowTabButton        bool  `gorm:"column:ShowTabButton"`
	ShowReadButton       bool  `gorm:"column:ShowReadButton"`
	ShowLinkButton       bool  `gorm:"column:ShowLinkButton"`
	ShowBookmarkButton   bool  `gorm:"column:ShowBookmarkButton"`
}

func (Settings) TableName() string {
	return "settings"
}

/* Twitter Models
============================================================================= */
type TwitterNews struct {
	Id          int64  `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	TweetId     int64  `gorm:"column:TweetId;index"`
	UserId      int64  `gorm:"column:UserId;index"`
	SourceId    int64  `gorm:"column:SourceId;index"`
	CreatedAt   int64  `gorm:"column:CreatedAt"`
	Text        string `gorm:"column:Text"`
	ExpandedUrl string `gorm:"column:ExpandedUrl"`
	Image       string `gorm:"column:Image"`
}

func (TwitterNews) TableName() string {
	return "twitternews"
}

type TwitterSource struct {
	Id         int64  `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	UserId     int64  `gorm:"column:UserId;index"`
	Name       string `gorm:"column:Name"`
	ScreenName string `gorm:"column:ScreenName"`
	Url        string `gorm:"column:Url"`
	Image      string `gorm:"column:Image"`
}

func (TwitterSource) TableName() string {
	return "twittersource"
}
