package dao

type Config struct {
	Driver           string `json:"driver"`
	ConnectionString string `json:"connection_string"`
	ClientSecret     string `json:"client_secret"`
	ApiKey           string `json:"api_key"`
	UpdateMinutes    int    `json:"update_minutes"`
}

type TwitterFriends struct {
	Users []TwitterUser `json:"users"`
}

type TwitterUser struct {
	ScreenName string `json:"screen_name"`
}

type TweetJson struct {
	Id       int64             `json:"id"`
	Text     string            `json:"text"`
	Source   TweetSourceJson   `json:"user"`
	Entities TweetEntitiesJson `json:"entities"`
}

type TweetSourceJson struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Url        string `json:"url"`
	Image      string `json:"profile_image_url_https"`
}

type TweetEntitiesJson struct {
	Urls  []TweetEntitiesUrlsJson  `json:"urls"`
	Media []TweetEntitiesMediaJson `json:"media"`
}

type TweetEntitiesUrlsJson struct {
	ExpandedUrl string `json:"expanded_url"`
}

type TweetEntitiesMediaJson struct {
	ExpandedUrl   string `json:"expanded_url"`
	MediaUrlHttps string `json:"media_url_https"`
}

func (t *TweetJson) ToDbSource() TwitterSource {
	source := TwitterSource{
		Name:       t.Source.Name,
		ScreenName: t.Source.ScreenName,
		Url:        t.Source.Url,
		Image:      t.Source.Image,
	}

	return source
}

func (t *TweetJson) ToDbNews() TwitterNews {
	news := TwitterNews{
		Text:    t.Text,
		TweetId: t.Id,
	}
	for _, urlsJson := range t.Entities.Urls {
		if urlsJson.ExpandedUrl != "" {
			news.ExpandedUrl = urlsJson.ExpandedUrl
			break
		}
	}
	for _, media := range t.Entities.Media {
		if media.MediaUrlHttps != "" {
			news.Image = media.MediaUrlHttps
			break
		}
	}

	return news
}
