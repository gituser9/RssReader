package main

import (
	"model"
	"flag"
	"io/ioutil"
	"os"
	"encoding/json"
	"time"
)

var conf model.Config

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

	if *pathPtr == defaultConfigPath {
		currentDir, _ := os.Getwd()
		conf = model.Config{FilePath: currentDir + string(os.PathSeparator) + "cfg.json"}
	} else {
		conf = model.Config{FilePath: *pathPtr}
	}

	// set default values
	conf.UpdateMinutes = 30

	json.Unmarshal(bytes, &conf)
}

func main() {
	updater := CreateUpdater(&conf)
	cleaner := CreateCleaner(&conf)
	updateTime := time.Duration(conf.UpdateMinutes) * time.Minute
	updateTimer := time.NewTicker(updateTime).C
	weekTimer := time.NewTicker(time.Hour * 168).C // week

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