package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"newshub/models"
	"newshub/services"

	"github.com/dgrijalva/jwt-go"
)

type UserController struct {
	service       *services.UserService
	config        *models.Config
	tokenLifeTime time.Duration
}

func NewUserCtrl(cfg *models.Config) *UserController {
	ctrl := new(UserController)
	ctrl.config = cfg
	ctrl.tokenLifeTime = 1 * time.Hour
	ctrl.service = services.NewUserService(cfg)

	return ctrl
}

func (ctrl *UserController) Auth(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	var login, password string
	ok := false

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if login, ok = data["username"]; !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if password, ok = data["password"]; !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	user := ctrl.service.Auth(login, password)

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonData, _ := json.Marshal(ctrl.createAuthData(user.Id))
	w.Write(jsonData)
}

func (ctrl *UserController) Registration(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	var login, password string
	ok := false

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if login, ok = data["username"]; !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if password, ok = data["password"]; !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	user, err := ctrl.service.Register(login, password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(ctrl.createAuthData(user.Id))
	w.Write(jsonData)
}

func (ctrl *UserController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	var token string
	ok := false

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token, ok = data["token"]; !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	claims := models.JwtClaims{}
	_, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(ctrl.config.JwtSign), nil
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonData, _ := json.Marshal(ctrl.createAuthData(claims.Id))
	w.Write(jsonData)
}

func (ctrl *UserController) GetUserSettings(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	settingsObj := services.SettingsService{}
	settingService := settingsObj.Init(ctrl.config)
	settings := settingService.Get(id)
	user := ctrl.service.GetUser(id)

	/*if user.VkNewsEnabled && len(user.VkPassword) > 0 {
		user.VkPassword = decryptVkPassword(user.VkPassword)
	}*/

	result := models.SettingsData{
		VkNewsEnabled:        settings.VkNewsEnabled,
		MarkSameRead:         settings.MarkSameRead,
		RssEnabled:           settings.RssEnabled,
		ShowPreviewButton:    settings.ShowPreviewButton,
		ShowReadButton:       settings.ShowReadButton,
		ShowTabButton:        settings.ShowTabButton,
		UnreadOnly:           settings.UnreadOnly,
		VkLogin:              user.VkLogin,
		VkPassword:           user.VkPassword,
		UserId:               id,
		TwitterEnabled:       settings.TwitterEnabled,
		TwitterName:          user.TwitterScreenName,
		TwitterSimpleVersion: settings.TwitterSimpleVersion,
		ShowLinkButton:       settings.ShowLinkButton,
		ShowBookmarkButton:   settings.ShowBookmarkButton,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (ctrl *UserController) SaveSettings(w http.ResponseWriter, r *http.Request) {
	settings := models.Settings{}
	body, err := ioutil.ReadAll(r.Body)
	claims := getClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &settings); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if claims.Id != settings.UserId {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	settingsObject := services.SettingsService{}
	settingService := settingsObject.Init(ctrl.config)
	settingService.Update(settings)
}

func (ctrl *UserController) createToken(id int64, duration time.Duration) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtClaims{
		Exp: time.Now().Add(duration).Unix(),
		Id:  id,
	})
	tokenString, _ := token.SignedString([]byte(ctrl.config.JwtSign))

	return tokenString
}

func (ctrl *UserController) createAuthData(id int64) *models.AuthData {
	return &models.AuthData{
		Token:        ctrl.createToken(id, ctrl.tokenLifeTime),
		RefreshToken: ctrl.createToken(id, 336*time.Hour),
	}
}
