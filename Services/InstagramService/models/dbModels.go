package dbModels

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
