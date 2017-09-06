package main

import (
	"flag"
	"io/ioutil"
	"encoding/json"
    "time"
    "os"

    "./models"
    "./services"
    "log"
)

var cfgPath string
var cfg *models.Configuration
const defaultConfigPath string = "./twitterConfig.json"

func readConfig() {
    bytes, err := ioutil.ReadFile(cfgPath)

    if err != nil {
        panic("Read config file error")
    }

    // set default values
    cfg = new(models.Configuration)
    cfg.UpdateMinutes = 30

    json.Unmarshal(bytes, cfg)

    if len(cfg.Driver) == 0 || len(cfg.ConnectionString) == 0 {
        panic("DB connection configuration failed")
    }
    if len(cfg.AccessToken) == 0 {
        panic("No access token")
    }
}

func init() {
	pathPtr := flag.String("config", defaultConfigPath, "Path for configuration file")
	flag.Parse()

	if *pathPtr == "" {
		panic("No config path")
	}
    if _, err := os.Stat(*pathPtr); os.IsNotExist(err) {
        panic("Not config file exists")
    }

    cfgPath = *pathPtr
    readConfig()
}

func main() {
	// get services
    netService := service.InitNetService(cfg)
    instagramService := service.InitInstagramService(cfg)
    userService := service.InitUserService(cfg)

    // create update timer
    updateTime := time.Duration(cfg.UpdateMinutes) * time.Minute
    updateTimer := time.NewTicker(updateTime).C

    // update with start
    for _, user := range userService.GetUsers() {
        log.Println(user.InstagramName)
        instagramService.Update(netService.GetData(user.InstagramName))
    }

    // run timer
    for {
       select {
       case <-updateTimer:
           for _, user := range userService.GetUsers() {
                instagramService.Update(netService.GetData(user.InstagramName))
           }
           // re-read config file
           readConfig()

           // update timer
           updateTime = time.Duration(cfg.UpdateMinutes) * time.Minute
       }
    }
}
