package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"model"
	"time"
)

var conf *model.Config

const defaultConfigPath = "./cfg.json"

func init() {
	// read config file
	pathPtr := flag.String("config", defaultConfigPath, "Path for configuration file")
	flag.Parse()

	if *pathPtr == "" {
		panic("No config path")
	}

	bytes, err := ioutil.ReadFile(*pathPtr)

	if err != nil {
		panic("Read config file error")
	}

	// set default values
	conf.UpdateMinutes = 30

	json.Unmarshal(bytes, conf)
}

func main() {
	updater := CreateUpdater(conf)
	cleaner := CreateCleaner(conf)

	updateTimer := time.Tick(time.Duration(conf.UpdateMinutes) * time.Minute)
	weekTimer := time.Tick(time.Hour * 168) // week

	go updater.Update()
	go cleaner.Clean()

	for {
		select {
		case <-updateTimer:
			updater.Update()
		case <-weekTimer:
			cleaner.Clean()
		}
	}
}
