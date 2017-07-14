package models

type Configuration struct {
	Driver           string `json:"driver"`
	ConnectionString string `json:"connection_string"`
	UpdateMinutes    int    `json:"update_minutes"`
}
