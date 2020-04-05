package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"newshub/models"
	"newshub/services"

	"github.com/gorilla/mux"
)

// RssController - request handlers
type RssController struct {
	service *services.RssService
	config  *models.Config
}

func NewRssCtrl(cfg *models.Config) *RssController {
	ctrl := new(RssController)
	ctrl.config = cfg
	ctrl.service = new(services.RssService).Init(cfg)

	return ctrl
}

// GetAll - get feed list
func (ctrl *RssController) GetAll(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	feeds := ctrl.service.GetRss(claims.Id)

	json.NewEncoder(w).Encode(feeds)
}

// GetArticles - get articles for feed
func (ctrl *RssController) GetArticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["feed_id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := getClaims(r)
	page, _ := strconv.Atoi(r.FormValue("page"))
	feed := ctrl.service.GetArticles(id, claims.Id, page)

	json.NewEncoder(w).Encode(feed)
}

// GetArticle - get one article
func (ctrl *RssController) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := getClaims(r)
	article := ctrl.service.GetArticle(id, claims.Id)

	if article == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// GetSettings - get app settings
/*func (ctrl *RssController) GetSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ctrl.config)
}*/

// AddFeed - add feed
func (ctrl *RssController) AddFeed(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	filters := models.RssFilters{}

	if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if filters.Url == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctrl.service.AddFeed(filters.Url, claims.Id)
	ctrl.GetAll(w, r)
}

// Delete - delete feed
func (ctrl *RssController) Delete(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctrl.service.Delete(id, claims.Id)
	ctrl.GetAll(w, r)
}

// UpdateAll - get new articles for all feeds
/*func (ctrl *RssController) UpdateAll(w http.ResponseWriter, r *http.Request) {
	ctrl.service.UpdateAllFeeds()
	ctrl.GetAll(w, r)
}*/

// UploadOpml - upload, parse OPML and update feeds
func (ctrl *RssController) UploadOpml(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	file, _, err := r.FormFile("file")

	if err != nil || file == nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctrl.service.Import(data, claims.Id)
	ctrl.GetAll(w, r)
}

func (ctrl *RssController) CreateOpml(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	ctrl.service.Export(claims.Id)
}

// SetNewFeedName - set new feed name
func (ctrl *RssController) SetNewFeedName(w http.ResponseWriter, r *http.Request) {
	feeds := models.FeedUpdateData{}

	if err := json.NewDecoder(r.Body).Decode(&feeds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := getClaims(r)
	ctrl.service.SetNewName(feeds, claims.Id)
	ctrl.GetAll(w, r)
}

// SetBookmark - set article is bookmark
func (ctrl *RssController) ToggleBookmark(w http.ResponseWriter, r *http.Request) {
	/*data := postClientData(r)
	ctrl.service.ToggleBookmark(data.ArticleId, data.IsBookmark)

	w.Header().Set("Content-Type", "application/json")
	success := make(map[string]bool, 1) // todo: to structure (field)
	success["success"] = true
	json.NewEncoder(w).Encode(success)*/
}

func (ctrl *RssController) GetBookmarks(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	page, err := strconv.Atoi(r.FormValue("page"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	articles := ctrl.service.GetBookmarks(page, claims.Id)

	json.NewEncoder(w).Encode(articles)
}

/*func (ctrl *RssController) MarkAllRead(w http.ResponseWriter, r *http.Request) {
	id64, _ := strconv.ParseUint(r.FormValue("id"), 10, 32)
	userId, _ := strconv.ParseUint(r.FormValue("userId"), 10, 32)
	id := uint(id64)
	ctrl.service.MarkReadAll(id)
	feed := ctrl.service.GetArticles(id, uint(userId), 1)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}*/

// fixme: remove
func (ctrl *RssController) ToggleUnread(w http.ResponseWriter, r *http.Request) {
	//data := postClientData(r)
	//ctrl.service.AppSettings.UnreadOnly = data.IsUnread // todo: auth and user setiings
}

func (ctrl *RssController) Search(w http.ResponseWriter, r *http.Request) {
	searchString := r.FormValue("searchString")
	isBookmark, _ := strconv.ParseBool(r.FormValue("isBookmark"))
	feedId, _ := strconv.ParseInt(r.FormValue("feedId"), 10, 64)
	claims := getClaims(r)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ctrl.service.Search(searchString, isBookmark, feedId, claims.Id))
}

func (ctrl *RssController) ToggleAsRead(w http.ResponseWriter, r *http.Request) {
	/*data := postClientData(r)

	ctrl.service.ToggleAsRead(data.ArticleId, data.IsRead)
	articles := ctrl.service.GetArticles(uint(data.FeedId), uint(data.UserId), data.Page)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)*/
}

func (ctrl *RssController) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := models.ArticlesUpdateData{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if id != data.ArticleId {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	claims := getClaims(r)
	ctrl.service.ArticleUpdate(claims.Id, data)
}
