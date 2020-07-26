package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"newshub-server/controllers"
	"newshub-server/middleware"
	"newshub-server/models"
	"newshub-server/services"

	"github.com/gorilla/mux"
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
	conf.OPMLPath, _ = os.Getwd()
	conf.PageSize = 20

	json.Unmarshal(bytes, &conf)

	services.Setup(conf)
}

func createRouter() http.Handler {
	rssCtrl := controllers.NewRssCtrl(&conf)
	userCtrl := controllers.NewUserCtrl(&conf)
	vkCtrl := controllers.NewVkCtrl(&conf)
	twitterCtrl := controllers.NewTwitterCtrl(&conf)
	router := mux.NewRouter()
	router.StrictSlash(true)

	// rss
	router.HandleFunc("/rss", rssCtrl.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/rss", rssCtrl.AddFeed).Methods(http.MethodPost)
	router.HandleFunc("/rss/{id}", rssCtrl.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/rss/{id}", rssCtrl.SetNewFeedName).Methods(http.MethodPut)
	router.HandleFunc("/rss/search", rssCtrl.Search).Methods(http.MethodGet)
	router.HandleFunc("/rss/opml", rssCtrl.UploadOpml).Methods(http.MethodPost)
	router.HandleFunc("/rss/opml", rssCtrl.CreateOpml).Methods(http.MethodGet)

	// articles
	router.HandleFunc("/rss/{feed_id}/articles", rssCtrl.GetArticles).Methods(http.MethodGet)
	router.HandleFunc("/rss/{feed_id}/articles/{id}", rssCtrl.GetArticle).Methods(http.MethodGet)
	router.HandleFunc("/rss/{feed_id}/articles/{id}", rssCtrl.UpdateArticle).Methods(http.MethodPut)
	router.HandleFunc("/rss/articles/bookmarks", rssCtrl.GetBookmarks)

	// user
	router.HandleFunc("/auth", userCtrl.Auth).Methods(http.MethodPost)
	router.HandleFunc("/registration", userCtrl.Registration).Methods(http.MethodPost)
	router.HandleFunc("/users/settings", userCtrl.GetUserSettings).Methods(http.MethodGet)
	router.HandleFunc("/users/settings", userCtrl.SaveSettings).Methods(http.MethodPut)
	router.HandleFunc("/users/refresh", userCtrl.RefreshToken).Methods(http.MethodPut)

	// vk
	router.HandleFunc("/vk", vkCtrl.GetPageData)
	router.HandleFunc("/vk/news", vkCtrl.GetNews)
	router.HandleFunc("/vk/search", vkCtrl.Search)

	// twitter
	router.HandleFunc("/twitter", twitterCtrl.GetPageData)
	router.HandleFunc("/twitter/news", twitterCtrl.GetNews)
	router.HandleFunc("/twitter/sources", twitterCtrl.GetSources)
	router.HandleFunc("/twitter/search", twitterCtrl.Search)

	// todo: client
	// static
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/"))))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dist/index.html")
	})

	// middleware
	amw := middleware.AuthenticationMiddleware{}
	amw.Populate(conf)

	router.Use(amw.Middleware)

	return router
}

func main() {
	// todo: websocket for update feed list
	controllers.Config = conf

	router := createRouter()
	log.Println("server start on", conf.Address)

	if err := http.ListenAndServe(conf.Address, router); err != nil {
		panic("start server error: " + err.Error())
	}
}
