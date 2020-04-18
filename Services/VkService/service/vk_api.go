package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"newshub-vk-service/model"
)

func getVkNews(login string, password string) ([]model.VkNews, error) {
	var news []model.VkNews
	token, err := getVkToken(login, password)

	if err != nil {
		return nil, err
	}

	requestUrl := "https://api.vk.com/method/newsfeed.get?filters=post&access_token=" + token
	response, err := http.Get(requestUrl)

	if err != nil {
		log.Println("get news error:", err)
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println("read news error:", err)
		return nil, err
	}

	log.Println(string(body))

	return news, nil
}

func getVkToken(login string, password string) (string, error) {
	requestUrl := fmt.Sprintf(
		"https://oauth.vk.com/token?grant_type=password&client_id=%d&client_secret=%s&username=%s&password=%s",
		"7404672", "b3e08548b3e08548b3e0854890b39079c8bb3e0b3e08548ed728525d785ad7f2301ab6b", login, password,
	)
	response, err := http.Post(requestUrl, "", nil)

	if err != nil {
		log.Println("get token error:", err)
		return "", err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println("read token error:", err)
		return "", err
	}

	return string(body), nil
}
