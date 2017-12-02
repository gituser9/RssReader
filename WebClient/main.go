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
	rssCtrl := new(controllers.RssController).Init(&conf)
	userCtrl := new(controllers.UserController).Init(&conf)
	vkCtrl := new(controllers.VkController).Init(&conf)
	twitterCtrl := new(controllers.TwitterController).Init(&conf)

	// todo: websocket for update feed list
	// todo: gorilla mux for REST API
	// todo: gorilla sessions

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/"))))
	http.HandleFunc("/", rssCtrl.Index)

	// rss
	http.HandleFunc("/get-all", rssCtrl.GetAll)
	http.HandleFunc("/get-articles", rssCtrl.GetArticles)
	http.HandleFunc("/get-article", rssCtrl.GetArticle)
	http.HandleFunc("/add-article", rssCtrl.AddFeed)
	http.HandleFunc("/delete", rssCtrl.Delete)
	http.HandleFunc("/set-new-name", rssCtrl.SetNewFeedName)
	http.HandleFunc("/update-all", rssCtrl.UpdateAll)
	http.HandleFunc("/upload-opml", rssCtrl.UploadOpml)
	http.HandleFunc("/toggle-bookmark", rssCtrl.ToggleBookmark)
	http.HandleFunc("/get-bookmarks", rssCtrl.GetBookmarks)
	http.HandleFunc("/mark-read-all", rssCtrl.MarkAllRead)
	http.HandleFunc("/create-opml", rssCtrl.CreateOpml)
	http.HandleFunc("/toggle-unread", rssCtrl.ToggleUnread)
	http.HandleFunc("/search", rssCtrl.Search)
	http.HandleFunc("/toggle-as-read", rssCtrl.ToggleAsRead)

	// user
	http.HandleFunc("/auth", userCtrl.Auth)
	http.HandleFunc("/registration", userCtrl.Registration)
	http.HandleFunc("/get-settings", userCtrl.GetUserSettings)
	http.HandleFunc("/set-settings", userCtrl.SaveSettings)

	// vk
	http.HandleFunc("/get-vk-page", vkCtrl.GetPageData)
	http.HandleFunc("/get-vk-news", vkCtrl.GetNews)
	http.HandleFunc("/get-vk-news-by-filters", vkCtrl.GetByFilters)
	http.HandleFunc("/search-vk-news", vkCtrl.Search)

	// twitter
	http.HandleFunc("/get-twitter-page", twitterCtrl.GetPageData)
	http.HandleFunc("/get-twitter-news", twitterCtrl.GetNews)
	http.HandleFunc("/get-twitter-news-by-filters", twitterCtrl.GetByFilters)
	http.HandleFunc("/search-twitter-news", twitterCtrl.Search)

	log.Println("server start on port " + strconv.Itoa(conf.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(conf.Port), nil)

	if err != nil {
		log.Println("Start rror: ", err.Error())
	}
}
