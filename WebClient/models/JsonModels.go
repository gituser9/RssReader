package models

// Config - app config, create from config file
type Config struct {
	Driver           string `json:"driver"`
	ConnectionString string `json:"connection_string"`
	OPMLPath         string `json:"opml_path"`
	UnreadOnly       bool   `json:"unread_only"`
	UpdateMinutes    int    `json:"update_minutes"`
	PageSize         int    `json:"page_size"`
	DbBackupPath     string `json:"db_backup_path"`
	FilePath         string
	Port             int `json:"port"`
	DownLoadThreads  int `json:"download_threads"`
}

type ClientData struct {
	FeedId       uint   `json:"feedId"`
	ArticleId    uint   `json:"articleId"`
	UserId       uint   `json:"userId"`
	Page         int    `json:"page"`
	SearchString string `json:"searchString"`
	Url          string `json:"url"`
	Name         string `json:"name"`
	IsRead       bool   `json:"isRead"`
	IsBookmark   bool   `json:"isBookmark"`
	IsUnread     bool   `json:"isUnread"`
}

type AuthData struct {
	Id       uint   `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

type SettingsData struct {
	UserId            uint   `json:"UserId"`
	VkLogin           string `json:"VkLogin"`
	VkPassword        string `json:"VkPassword"`
	TwitterName       string `json:"TwitterName"`
	VkNewsEnabled     bool   `json:"VkNewsEnabled"`
	TwitterEnabled    bool   `json:"TwitterEnabled"`
	TwitterSimpleVersion    bool   `json:"TwitterSimpleVersion"`
	MarkSameRead      bool   `json:"MarkSameRead"`
	UnreadOnly        bool   `json:"UnreadOnly"`
	RssEnabled        bool   `json:"RssEnabled"`
	ShowPreviewButton bool   `json:"ShowPreviewButton"`
	ShowTabButton     bool   `json:"ShowTabButton"`
	ShowReadButton    bool   `json:"ShowReadButton"`
	ShowLinkButton    bool   `json:"ShowLinkButton"`
	ShowBookmarkButton    bool   `json:"ShowBookmarkButton"`
}

type VkData struct {
	GroupId      int    `json:"GroupId"`
	SearchString string `json:"SearchString"`
}

type TwitterData struct {
	SourceId int `json:"SourceId"`
	SearchString string `json:"SearchString"`
}
