package models

type Users struct {
	Id                int   `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name              string `gorm:"column:Name"`
	Password          string `gorm:"column:Password"`
	VkLogin           string `gorm:"column:VkLogin"`
	VkPassword        string `gorm:"column:VkPassword"`
	TwitterScreenName string `gorm:"column:TwitterScreenName"`
	InstagramName string `gorm:"column:InstagramName"`
	VkNewsEnabled     bool   `gorm:"column:VkNewsEnabled"`
	Settings          Settings
}

func (Users) TableName() string {
	return "users"
}

type Settings struct {
	Id                int `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	UserId            int `gorm:"column:UserId;index"`
	UnreadOnly        bool `gorm:"column:UnreadOnly"`
	MarkSameRead      bool `gorm:"column:MarkSameRead"`
	RssEnabled        bool `gorm:"column:RssEnabled"`
	VkNewsEnabled     bool `gorm:"column:VkNewsEnabled"`
	TwitterEnabled    bool `gorm:"column:TwitterEnabled"`
	InstagramEnabled    bool `gorm:"column:InstagramEnabled"`
	ShowPreviewButton bool `gorm:"column:ShowPreviewButton"`
	ShowTabButton     bool `gorm:"column:ShowTabButton"`
	ShowReadButton    bool `gorm:"column:ShowReadButton"`
}

func (Settings) TableName() string {
	return "settings"
}

type InstagramNews struct {
	Id     int
	UserId int
	Url    string
}

func (InstagramNews) TableName() string {
    return "instagramnews"
}
