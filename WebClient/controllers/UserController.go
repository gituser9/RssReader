package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	"../models"
	"../services"
)

type UserController struct {
	service *services.UserService
	config  *models.Config
}

// Init - init controller
func (ctrl *UserController) Init(config *models.Config) *UserController {
	service := new(services.UserService).Init(config)

	return &UserController{service: service, config: config}
}

func (ctrl *UserController) Auth(w http.ResponseWriter, r *http.Request) {
	authData := postUserData(r)
	user := ctrl.service.Auth(authData.Name, authData.Password)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (ctrl *UserController) Registration(w http.ResponseWriter, r *http.Request) {
	authData := postUserData(r)
	result := ctrl.service.Register(authData.Name, authData.Password)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (ctrl *UserController) GetUserSettings(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(r.URL.Query().Get("id"))

	settingService := services.SettingsService.Init(ctrl.config)
	settings := settingService.Get(uint(userId))
	user := ctrl.service.GetUser(uint(userId))
	result := models.SettingsData{
		VkNewsEnabled: settings.VkNewsEnabled,
		MarkSameRead: settings.MarkSameRead,
		RssEnabled: settings.RssEnabled,
		ShowPreviewButton: settings.ShowPreviewButton,
		ShowReadButton: settings.ShowReadButton,
		ShowTabButton: settings.ShowTabButton,
		UnreadOnly: settings.UnreadOnly,
		VkLogin: user.VkLogin,
		VkPassword: user.VkPassword,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (ctrl *UserController) SaveSettings(w http.ResponseWriter, r *http.Request) {
	settingsData := postSettingsData(r)
	settings := models.Settings{
		MarkSameRead:      settingsData.MarkSameRead,
		RssEnabled:        settingsData.RssEnabled,
		UnreadOnly:        settingsData.UnreadOnly,
		VkNewsEnabled:     settingsData.VkNewsEnabled,
		ShowPreviewButton: settingsData.ShowPreviewButton,
		ShowReadButton:    settingsData.ShowReadButton,
		ShowTabButton:     settingsData.ShowTabButton,
	}
	settingService := services.SettingsService.Init(ctrl.config)
	existingSettings := settingService.Get(settingsData.UserId)
	existingSettings = settings

	settingService.Update(existingSettings)

	if settingsData.VkNewsEnabled {
		user := ctrl.service.GetUser(settingsData.UserId)
		user.VkLogin = settingsData.VkLogin
		user.VkPassword = settingsData.VkPassword

		ctrl.service.Update(user)
	}

	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(result)
}

func postUserData(r *http.Request) models.AuthData {
	result := new(models.AuthData)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)

	if err != nil {
		log.Println("decode err: ", err.Error())
	}

	return *result
}

func postSettingsData(r *http.Request) models.SettingsData {
	result := new(models.SettingsData)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)

	if err != nil {
		log.Println("decode err: ", err.Error())
	}

	return *result
}
