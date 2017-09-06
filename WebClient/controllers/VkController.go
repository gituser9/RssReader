package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"../models"
	"../services"
)

type VkController struct {
	service *services.VkService
	config  *models.Config
}

// Init - init controller
func (ctrl *VkController) Init(config *models.Config) *VkController {
	service := new(services.VkService).Init(config)

	return &VkController{service: service, config: config}
}

func (ctrl *VkController) GetPageData(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	pageData := models.VkPageData{
		News:   ctrl.service.GetNews(id, 1),
		Groups: ctrl.service.GetAllGroups(id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pageData)
}

func (ctrl *VkController) GetNews(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	news := ctrl.service.GetNews(id, page)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func (ctrl *VkController) GetByFilters(w http.ResponseWriter, r *http.Request) {
	data := postVkData(r)
	news := ctrl.service.GetNewsByFilters(data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func (ctrl *VkController) Search(w http.ResponseWriter, r *http.Request) {
	data := postVkData(r)
	news := ctrl.service.Search(data.SearchString, data.GroupId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

/*==============================================================================
	Private
==============================================================================*/

func postVkData(r *http.Request) models.VkData {
	result := new(models.VkData)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)

	if err != nil {
		log.Println("decode err: ", err.Error())
	}

	return *result
}
