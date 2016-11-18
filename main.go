package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"./controllers"
	"./models"
	"./services"
)

var conf models.Config

const defaultConfigPath = "./cfg.json"

func startTimers(config *models.Config) {
	// todo: DB backup timer (24 hours)
	service := new(services.RssService).Init(config)

	updateTime := time.Duration(service.AppSettings.UpdateMinutes) * time.Minute
	updateTimer := time.NewTicker(updateTime).C
	weekTimer := time.NewTicker(time.Hour * 168).C // week

	// on start
	//service.CleanOldArticles()
	go service.UpdateAllFeeds()

	for {
		select {
		case <-updateTimer:
			service.UpdateAllFeeds()
		case <-weekTimer:
			service.CleanOldArticles()
			service.Export()
			service.Backup()
		}
	}
}

func init() {
	// read config file
	pathPtr := flag.String("config", defaultConfigPath, "Path for configuration file")

	if *pathPtr == "" {
		panic("No config path")
	}

	bytes, err := ioutil.ReadFile(*pathPtr)

	if err != nil {
		panic("Read config file error")
	}

	if *pathPtr == defaultConfigPath {
		currentDir, _ := os.Getwd()
		conf = models.Config{FilePath: currentDir + string(os.PathSeparator) + "cfg.json"}
	} else {
		conf = models.Config{FilePath: *pathPtr}
	}

	// set default values
	conf.UpdateMinutes = 30
	conf.DownLoadThreads = 4
	conf.OPMLPath, _ = os.Getwd()
	conf.PageSize = 20

	json.Unmarshal(bytes, &conf)
}

func main() {
	controller := new(controllers.RssController).Init(&conf)

	go startTimers(&conf)

	// todo: websocket for update feed list
	// todo: gorilla mux for REST API

	log.Println("server start")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/"))))
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/get-all", controller.GetAll)
	http.HandleFunc("/get-articles", controller.GetArticles)
	http.HandleFunc("/get-article", controller.GetArticle)
	http.HandleFunc("/add-article", controller.AddFeed)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/set-new-name", controller.SetNewFeedName)
	http.HandleFunc("/update-all", controller.UpdateAll)
	http.HandleFunc("/upload-opml", controller.UploadOpml)
	http.HandleFunc("/toggle-bookmark", controller.ToggleBookmark)
	http.HandleFunc("/get-bookmarks", controller.GetBookmarks)
	http.HandleFunc("/mark-read-all", controller.MarkAllRead)
	http.HandleFunc("/create-opml", controller.CreateOpml)
	http.HandleFunc("/toggle-unread", controller.ToggleUnread)
	http.HandleFunc("/get-settings", controller.GetAppSettings)
	http.HandleFunc("/search", controller.Search)
	http.HandleFunc("/toggle-as-read", controller.ToggleAsRead)
	http.ListenAndServe(":"+strconv.Itoa(conf.Port), nil)
}
