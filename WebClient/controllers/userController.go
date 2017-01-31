package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"../models"
	"../services"
	"strconv"
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
	user := ctrl.service.GetUser(uint(userId))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (ctrl *UserController) SaveSettings(w http.ResponseWriter, r *http.Request) {
	//settingsData := postSettingsData(r)

	// get JSON

	//settings := models.Settings{}
	//_ := models.Users{Id: settingsData.UserId}

	// set settings
	//service := services.SettingsService.Init(ctrl.config)
	//service.Update(settings)

	//w.Header().Set("Content-Type", "application/json")
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
