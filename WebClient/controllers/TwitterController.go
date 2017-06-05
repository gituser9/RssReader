package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../models"
	"../services"
)

type TwitterController struct {
	service *services.TwitterService
	config  *models.Config
}

// Init - init controller
func (ctrl *TwitterController) Init(config *models.Config) *TwitterController {
	service := new(services.TwitterService).Init(config)

	return &TwitterController{service: service, config: config}
}

func (ctrl *TwitterController) GetPageData(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	pageData := models.TwitterPageData{
		News:    ctrl.service.GetNews(id, 1),
		Sources: ctrl.service.GetAllSources(id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pageData)
}

func (ctrl *TwitterController) GetNews(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	news := ctrl.service.GetNews(id, page)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func (ctrl *TwitterController) GetByFilters(w http.ResponseWriter, r *http.Request) {
	data := postTwitterData(r)
	news := ctrl.service.GetNewsByFilters(data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

/*==============================================================================
	Private
==============================================================================*/

func postTwitterData(r *http.Request) models.TwitterData {
	result := new(models.TwitterData)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)

	if err != nil {
		log.Println("decode err: ", err.Error())
	}

	return *result
}
