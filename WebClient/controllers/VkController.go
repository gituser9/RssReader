package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		News:   ctrl.service.GetAllNews(id),
		Groups: ctrl.service.GetAllGroups(id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pageData)
}

func (ctrl *VkController) GetAll(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	news := ctrl.service.GetAllNews(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}
