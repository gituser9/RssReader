package service

import (
	"encoding/json"
	"net/http"
	"net/url"

    "../models"
    "log"
    "io/ioutil"
)

const apiUrl string = "https://api.instagram.com/v1/"

type NetService struct {
    cfg *models.Configuration
}

func (service *NetService) GetData(name string) string {
    if len(name) == 0 {
        return ""
    }

    log.Println(service.getUserId(name))

    return ""
}


func (service *NetService) getNews() {

}

func (service *NetService) getUserId(name string) string {
    urlParams := map[string]string {
        "q": name,
        "access_token": service.cfg.AccessToken,
    }
    response, err := http.Get(createUrl("users/search", urlParams))
    defer response.Body.Close()

    if err != nil {
        log.Println(name + " get data error: " + err.Error())
        return ""
    }

    jsonData, err := ioutil.ReadAll(response.Body)

    if err != nil {
        log.Println("Read user data error: " + err.Error())
    }

    userData := models.UserDataJson{}
    json.Unmarshal(jsonData, &userData)

    if len(userData.Data) > 0 {
        return userData.Data[0].Id
    } else {
        return ""
    }
}

func InitNetService(cfg *models.Configuration) *NetService {
    netService := new(NetService)
    netService.cfg = cfg

    return netService
}

func createUrl(path string, params map[string]string) string {
    uri, _ := url.Parse(apiUrl + path)
    query := uri.Query()

    for key, value := range params {
        query.Set(key, value)
    }

    uri.RawQuery = query.Encode()

    return uri.String()
}