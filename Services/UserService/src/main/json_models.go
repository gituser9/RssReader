package main

type Config struct {
	Driver           string `json:"driver"`
	ConnectionString string `json:"connection_string"`
	Port             int `json:"port"`
}