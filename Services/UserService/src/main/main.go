package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"strconv"

	"github.com/claygod/door"
)

var cfg Config

const defaultConfigPath = "./cfg.json"

func init() {
	// read config file
	pathPtr := flag.String("config", defaultConfigPath, "Path for configuration file")
	flag.Parse()

	if pathPtr == nil || *pathPtr == "" {
		panic("No config path")
	}

	bytes, err := ioutil.ReadFile(*pathPtr)

	if err != nil {
		panic("Read config file error")
	}

	// set default values
	cfg.Port = 8900

	json.Unmarshal(bytes, &cfg)
}

func main() {
	handler := CreateHttphandler(cfg)
	router := door.New()

	router.Add("/users/rss", handler.GetUsersIdWithRssEnabled)
	router.Start(":" + strconv.Itoa(cfg.Port))
}
