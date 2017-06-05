package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

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
	log.Println(authData)
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
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	userId := uint(id)

	settingsObj := services.SettingsService{}
	settingService := settingsObj.Init(ctrl.config)
	settings := settingService.Get(userId)
	user := ctrl.service.GetUser(userId)

	/*if user.VkNewsEnabled && len(user.VkPassword) > 0 {
		user.VkPassword = decryptVkPassword(user.VkPassword)
	}*/

	result := models.SettingsData{
		VkNewsEnabled:     settings.VkNewsEnabled,
		MarkSameRead:      settings.MarkSameRead,
		RssEnabled:        settings.RssEnabled,
		ShowPreviewButton: settings.ShowPreviewButton,
		ShowReadButton:    settings.ShowReadButton,
		ShowTabButton:     settings.ShowTabButton,
		UnreadOnly:        settings.UnreadOnly,
		VkLogin:           user.VkLogin,
		VkPassword:        user.VkPassword,
		UserId:            userId,
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
		UserId:            settingsData.UserId,
	}
	settingsObject := services.SettingsService{}
	settingService := settingsObject.Init(ctrl.config)
	settingService.Update(settings)

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

func decryptVkPassword(decryptedPassword string) string {
	//key := []byte("AES256Key-32Characters1234567890")
	key := []byte("AES128Key-16Char")
	ciphertext, _ := hex.DecodeString(decryptedPassword)
	nonce, _ := hex.DecodeString("37b8e8a308c354048d245f6d")
	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		return decryptedPassword
	}

	return string(plaintext)
}
