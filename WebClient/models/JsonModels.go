package models

import "github.com/dgrijalva/jwt-go"

// Config - app config, create from config file
type Config struct {
	Address          string `json:"address"`
	Driver           string `json:"driver"`
	ConnectionString string `json:"connection_string"`
	OPMLPath         string `json:"opml_path"`
	UnreadOnly       bool   `json:"unread_only"`
	UpdateMinutes    int    `json:"update_minutes"`
	PageSize         int    `json:"page_size"`
	DbBackupPath     string `json:"db_backup_path"`
	FilePath         string
	DownLoadThreads  int    `json:"download_threads"`
	JwtSign          string `json:"jwt_sign"`
}

type SettingsData struct {
	UserId               int64  `json:"UserId"`
	VkLogin              string `json:"VkLogin"`
	VkPassword           string `json:"VkPassword"`
	TwitterName          string `json:"TwitterName"`
	VkNewsEnabled        bool   `json:"VkNewsEnabled"`
	TwitterEnabled       bool   `json:"TwitterEnabled"`
	TwitterSimpleVersion bool   `json:"TwitterSimpleVersion"`
	MarkSameRead         bool   `json:"MarkSameRead"`
	UnreadOnly           bool   `json:"UnreadOnly"`
	RssEnabled           bool   `json:"RssEnabled"`
	ShowPreviewButton    bool   `json:"ShowPreviewButton"`
	ShowTabButton        bool   `json:"ShowTabButton"`
	ShowReadButton       bool   `json:"ShowReadButton"`
	ShowLinkButton       bool   `json:"ShowLinkButton"`
	ShowBookmarkButton   bool   `json:"ShowBookmarkButton"`
}

type VkData struct {
	GroupId      int64  `json:"GroupId"`
	SearchString string `json:"SearchString"`
}

type TwitterData struct {
	SourceId     int64  `json:"SourceId"`
	SearchString string `json:"SearchString"`
}

type RssFilters struct {
	Search      string `json:"search"`
	MarkAllRead bool   `json:"mark_all_read"`
	Name        string `json:"name"`
	Url         string `json:"url"`
}

type AuthData struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type JwtClaims struct {
	*jwt.MapClaims
	Id  int64
	Exp int64
}

type ArticlesUpdateData struct {
	ArticleId  int64 `json:"article_id"`
	IsRead     bool  `json:"is_read"`
	IsBookmark bool  `json:"is_bookmark"`
}

type FeedUpdateData struct {
	FeedId    int64  `json:"feed_id"`
	Name      string `json:"name"`
	IsReadAll bool   `json:"is_read_all"`
}
