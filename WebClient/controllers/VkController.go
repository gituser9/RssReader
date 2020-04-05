package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"newshub/models"
	"newshub/services"
)

type VkController struct {
	service *services.VkService
	config  *models.Config
}

func NewVkCtrl(cfg *models.Config) *VkController {
	ctrl := new(VkController)
	ctrl.config = cfg
	ctrl.service = services.NewVkService(cfg)

	return ctrl
}

func (ctrl *VkController) GetPageData(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	pageData := models.VkPageData{
		News:   ctrl.service.GetNews(claims.Id, 1, 0),
		Groups: ctrl.service.GetAllGroups(claims.Id),
	}

	json.NewEncoder(w).Encode(pageData)
}

func (ctrl *VkController) GetNews(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	groupId := int64(0)

	if r.FormValue("group_id") != "" {
		groupId, err = strconv.ParseInt(r.FormValue("group_id"), 10, 64)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	news := ctrl.service.GetNews(claims.Id, page, groupId)

	json.NewEncoder(w).Encode(news)
}

func (ctrl *VkController) Search(w http.ResponseWriter, r *http.Request) {
	claims := getClaims(r)
	groupId := int64(0)
	var err error

	if r.FormValue("group_id") != "" {
		groupId, err = strconv.ParseInt(r.FormValue("group_id"), 10, 64)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	news := ctrl.service.Search(r.FormValue("q"), groupId, claims.Id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func getVkFilters(r *http.Request) models.VkData {
	result := models.VkData{}

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		log.Println("decode err: ", err)
	}

	return result
}
