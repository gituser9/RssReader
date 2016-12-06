package controllers

import (
	"encoding/json"
	"log"
	"net/http"

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

func postUserData(r *http.Request) models.AuthData {
	result := new(models.AuthData)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)

	if err != nil {
		log.Println("decode err: ", err.Error())
	}

	return *result
}
