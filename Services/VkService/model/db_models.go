package model

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

type VkNews struct {
	Id        int64  `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	UserId    int64  `gorm:"column:UserId;index"`
	GroupId   int64  `gorm:"column:GroupId;index"`
	PostId    int64  `gorm:"column:PostId;index"`
	Timestamp int64  `gorm:"column:Timestamp"`
	Text      string `gorm:"column:Text"`
	Image     string `gorm:"column:Image"`
	Link      string `gorm:"column:Link"`
}

func (VkNews) TableName() string {
	return "vknews"
}

type VkGroup struct {
	Id         int64  `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Gid        int64  `gorm:"column:Gid;index"`
	UserId     int64  `gorm:"column:UserId;index"`
	Name       string `gorm:"column:Name"`
	LinkedName string `gorm:"column:LinkedName"`
	Image      string `gorm:"column:Image"`
}

func (VkGroup) TableName() string {
	return "vkgroups"
}

type Users struct {
	Id                int64    `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name              string   `gorm:"column:Name"`
	Password          string   `gorm:"column:Password" json:"-"`
	VkLogin           string   `gorm:"column:VkLogin"`
	VkPassword        string   `gorm:"column:VkPassword"`
	TwitterScreenName string   `gorm:"column:TwitterScreenName"`
	VkNewsEnabled     bool     `gorm:"column:VkNewsEnabled"`
	Settings          Settings `gorm:"ForeignKey:UserId"`
}

func (Users) TableName() string {
	return "users"
}
