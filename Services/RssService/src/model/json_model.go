package model

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