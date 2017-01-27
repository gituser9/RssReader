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
