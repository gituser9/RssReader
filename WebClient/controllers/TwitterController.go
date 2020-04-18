package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"newshub-server/models"
	"newshub-server/services"
)

type TwitterController struct {
	service *services.TwitterService
	config  *models.Config
}

func NewTwitterCtrl(cfg *models.Config) *TwitterController {
	ctrl := new(TwitterController)
	ctrl.config = cfg
	ctrl.service = services.NewTwitterService(cfg)

	return ctrl
}

func (ctrl *TwitterController) GetPageData(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	pageData := models.TwitterPageData{
		News:    ctrl.service.GetNews(claims.Id, 1, 0),
		Sources: ctrl.service.GetAllSources(claims.Id),
	}

	json.NewEncoder(w).Encode(pageData)
}

func (ctrl *TwitterController) GetNews(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	page, err := strconv.Atoi(r.FormValue("page"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	sourceId := int64(0)

	if r.FormValue("source_id") != "" {
		sourceId, err = strconv.ParseInt(r.FormValue("source_id"), 10, 64)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	news := ctrl.service.GetNews(claims.Id, page, sourceId)
	json.NewEncoder(w).Encode(news)
}

func (ctrl *TwitterController) GetSources(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	sources := ctrl.service.GetAllSources(claims.Id)
	json.NewEncoder(w).Encode(sources)
}

func (ctrl *TwitterController) Search(w http.ResponseWriter, r *http.Request) {
	sourceId := int64(0)
	var err error

	if r.FormValue("source_id") != "" {
		sourceId, err = strconv.ParseInt(r.FormValue("source_id"), 10, 64)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	claims := getClaims(r)
	news := ctrl.service.Search(r.FormValue("search_string"), sourceId, claims.Id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}
