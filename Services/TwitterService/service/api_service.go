package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"newshub-twitter-service/dao"
)

func getNews(name string) []dao.TweetJson {
	news := []dao.TweetJson{}
	token, tokenType := getToken()

	if token == "" || tokenType == "" {
		return news
	}

	friends := getFriends(name, tokenType, token)
	client := http.Client{}

	for _, friend := range friends {
		requestUrl := "https://api.twitter.com/1.1/statuses/user_timeline.json?include_rts=1&exclude_replies=1&screen_name=" + friend
		log.Println(requestUrl)

		req, _ := http.NewRequest(http.MethodGet, requestUrl, nil)
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", tokenType, token))

		response, err := client.Do(req)
		items := []dao.TweetJson{}

		if err != nil {
			log.Println("get news error:", err)
			continue
		}
		if err := json.NewDecoder(response.Body).Decode(&items); err != nil {
			log.Println("decode news error:", err)
			continue
		}

		news = append(news, items...)
	}

	return news
}

func getFriends(name, tokenType, token string) []string {
	names := []string{}
	requestUrl := fmt.Sprintf("https://api.twitter.com/1.1/friends/list.json?cursor=-1&screen_name=%s&skip_status=true&include_user_entities=false", name)
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, requestUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", tokenType, token))

	friends := dao.TwitterFriends{}
	res, err := client.Do(req)

	if err != nil {
		return names
	}
	if err := json.NewDecoder(res.Body).Decode(&friends); err != nil {
		log.Println("decode friends err:", err)
		return names
	}
	for _, friend := range friends.Users {
		names = append(names, friend.ScreenName)
	}

	return names
}

func getToken() (string, string) {
	authString := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.ApiKey, config.ClientSecret)))

	client := http.Client{}
	val := url.Values{
		"grant_type": []string{"client_credentials"},
	}
	req, _ := http.NewRequest(http.MethodPost, "https://api.twitter.com/oauth2/token", strings.NewReader(val.Encode()))
	req.Header.Set("Authorization", "Basic "+authString)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	response, err := client.Do(req)
	data := map[string]string{}

	if err != nil {
		log.Println("get token err:", err)
		return "", ""
	}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Println("decode token error:", err)
		return "", ""
	}

	return data["access_token"], data["token_type"]
}
