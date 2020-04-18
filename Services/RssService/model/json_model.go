package model

// Config - application configuration structure
type Config struct {
	Driver           string `json:"driver"`
	ConnectionString string `json:"connection_string"`
	UpdateMinutes    int    `json:"update_minutes"`
	ArticlesMaxCount int    `json:"articles_max_count"`
}
