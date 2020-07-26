package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"newshub-twitter-service/dao"
	"newshub-twitter-service/service"
)

var conf dao.Config

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

	if err := json.Unmarshal(bytes, &conf); err != nil {
		panic(err.Error())
	}
}

func main() {

	updateTimer := time.Tick(time.Duration(conf.UpdateMinutes) * time.Minute)
	weekTimer := time.Tick(time.Hour * 168) // week

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	service.Setup(conf)
	go service.Update()
	go service.Clean()

	for {
		select {
		case <-updateTimer:
			go service.Update()
		case <-weekTimer:
			go service.Clean()
		case <-sigs:
			service.Close()
			return
		}
	}
}
