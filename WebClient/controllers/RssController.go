package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../models"
	"../services"
)

// RssController - request handlers
type RssController struct {
	service *services.RssService
	config  *models.Config
}

// Init - init controller
func (ctrl *RssController) Init(config *models.Config) *RssController {
	service := new(services.RssService).Init(config)

	return &RssController{service: service, config: config}
}

// Index - return page
func (ctrl *RssController) Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dist/index.html")
}

// GetAll - get feed list
func (ctrl *RssController) GetAll(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	feeds := ctrl.service.GetRss(uint(id))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feeds)
}

// GetArticles - get articles for feed
func (ctrl *RssController) GetArticles(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
	feed := ctrl.service.GetArticles(uint(id), int(page))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}

// GetArticle - get one article
func (ctrl *RssController) GetArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	article := ctrl.service.GetArticle(uint(id))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// GetSettings - get app settings
func (ctrl *RssController) GetSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ctrl.config)
}

// AddFeed - add feed
func (ctrl *RssController) AddFeed(w http.ResponseWriter, r *http.Request) {
	data := postClientData(r)

	ctrl.service.AddFeed(data.Url, data.UserId)
	ctrl.GetAll(w, r)
}

// Delete - delete feed
func (ctrl *RssController) Delete(w http.ResponseWriter, r *http.Request) {
	data := postClientData(r)

	ctrl.service.Delete(data.FeedId)
	ctrl.GetAll(w, r)
}

// UpdateAll - get new articles for all feeds
func (ctrl *RssController) UpdateAll(w http.ResponseWriter, r *http.Request) {
	ctrl.service.UpdateAllFeeds()
	ctrl.GetAll(w, r)
}

// UploadOpml - upload, parse OPML and update feeds
func (ctrl *RssController) UploadOpml(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	id, _ := strconv.Atoi(r.FormValue("userId"))

	if err != nil || file == nil {
		log.Println(err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Println(err)
	}

	ctrl.service.Import(data, uint(id))
	ctrl.GetAll(w, r)
}

func (ctrl *RssController) CreateOpml(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	ctrl.service.Export(uint(id))
}

// SetNewFeedName - set new feed name
func (ctrl *RssController) SetNewFeedName(w http.ResponseWriter, r *http.Request) {
	data := postClientData(r)

	ctrl.service.SetNewName(data.Name, data.FeedId)
	ctrl.GetAll(w, r)
}

// SetBookmark - set article is bookmark
func (ctrl *RssController) ToggleBookmark(w http.ResponseWriter, r *http.Request) {
	data := postClientData(r)
	ctrl.service.ToggleBookmark(data.ArticleId, data.IsBookmark)

	w.Header().Set("Content-Type", "application/json")
	success := make(map[string]bool, 1) // todo: to structure (field)
	success["success"] = true
	json.NewEncoder(w).Encode(success)
}

func (ctrl *RssController) GetBookmarks(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
	articles := ctrl.service.GetBookmarks(page)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func (ctrl *RssController) MarkAllRead(w http.ResponseWriter, r *http.Request) {
	id64, _ := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	id := uint(id64)
	ctrl.service.MarkReadAll(id)
	feed := ctrl.service.GetArticles(id, 1)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}

// fixme: remove
func (ctrl *RssController) ToggleUnread(w http.ResponseWriter, r *http.Request) {
	data := postClientData(r)
	ctrl.service.AppSettings.UnreadOnly = data.IsUnread // todo: auth and user setiings
}

func (ctrl *RssController) GetAppSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ctrl.service.AppSettings)
}

func (ctrl *RssController) Search(w http.ResponseWriter, r *http.Request) {
	searchString := r.URL.Query().Get("searchString")
	isBookmark, _ := strconv.ParseBool(r.URL.Query().Get("isBookmark"))
	feedId, _ := strconv.Atoi(r.URL.Query().Get("feedId"))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ctrl.service.Search(searchString, isBookmark, uint(feedId)))
}

func (ctrl *RssController) ToggleAsRead(w http.ResponseWriter, r *http.Request) {
	data := postClientData(r)

	ctrl.service.ToggleAsRead(data.ArticleId, data.IsRead)
	articles := ctrl.service.GetArticles(uint(data.FeedId), data.Page)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

/*==============================================================================
	Private
==============================================================================*/

func postClientData(r *http.Request) models.ClientData {
	result := new(models.ClientData)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)

	if err != nil {
		log.Println("decode err: ", err.Error())
	}

	return *result
}
