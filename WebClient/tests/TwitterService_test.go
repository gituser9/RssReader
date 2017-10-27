package tests

import (
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"strconv"

	"../models"
	"../services"
)

var twitterNews = make([]models.TwitterNews, 10)

func mockDbForTwitter() *gorm.DB {
	db, _ = gorm.Open("sqlite3", dbName)
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Settings{})
	db.AutoMigrate(&models.TwitterSource{})
	db.AutoMigrate(&models.TwitterNews{})

	user = models.Users{
		Id: userId,
		Name:     "Test User",
		Password: "TestPassword",
		Settings: models.Settings{
			Id:           1,
			UnreadOnly:   false,
			UserId:       userId,
			RssEnabled:   true,
			MarkSameRead: true,
		},
	}
	db.Create(&user)

	for i := 1; i < 11; i++ {
		indexString := strconv.Itoa(i)
		twitterSource := models.TwitterSource{
			UserId: vkUserId,
			Name: "Twitter Source " + indexString,
			Url: "Url_" + indexString,
			Image: "image_" + indexString,
			ScreenName: "twitter_source_" + indexString,
		}
		db.Create(&twitterSource)

		for j := 1; j < 11; j++ {
			stringNum := strconv.Itoa(j + (i * 10))
			twitterItem := models.TwitterNews{
				Image: "Image " + stringNum,
				UserId: vkUserId,
				ExpandedUrl: "expanded_url_" + stringNum,
				SourceId: i,
				Text: "TwitterNews Text " + stringNum,
			}
			db.Create(&twitterItem)
		}
	}
	db.Find(&twitterNews)

	return db
}

func fakeInitTwitterService() *services.TwitterService {
	db := mockDbForTwitter()
	config := models.Config{
		PageSize:   20,
		UnreadOnly: false,
		OPMLPath:   ".",
	}
	service := services.TwitterService{}
	service.SetDb(db)
	service.SetConfig(&config)

	return &service
}

func TestGetTwitterNews(t *testing.T) {
	log.Println("Start TestGetTwitterNews")
	service := fakeInitTwitterService()
	defer os.Remove(dbName)

	news := service.GetNews(vkUserId, 1)

	if len(news) == 0 {
		t.Error("Twitter News is empty")
	}
}

func TestGetTwitterSources(t *testing.T) {
	log.Println("Start TestGetTwitterSources")
	service := fakeInitTwitterService()
	defer os.Remove(dbName)

	sources := service.GetAllSources(vkUserId)

	if len(sources) == 0 {
		t.Error("Sources is empty")
	}
}

func TestGetTwitterNewsByFilter(t *testing.T) {
	log.Println("Start TestGetTwitterNewsByFilter")
	service := fakeInitTwitterService()
	defer os.Remove(dbName)

	for _, item := range twitterNews {
	    filters := models.TwitterData{
	    	SourceId: item.SourceId,
		}
		news := service.GetNewsByFilters(filters)

		if len(news) == 0 {
			t.Error("News is empty")
			continue
		}
		for _, newsItem := range news {
			if newsItem.SourceId != filters.SourceId {
				t.Error("Filter by group error")
			}
		}

		filters.SearchString = item.Text
		filters.SourceId = 0
		news = service.GetNewsByFilters(filters)

		if len(news) == 0 {
			t.Error("Filter by string error")
		}
	}
}

// TODO: sources
func TestTwitterSearch(t *testing.T) {
	log.Println("Start TestTwitterSearch")
	service := fakeInitTwitterService()
	defer os.Remove(dbName)

	for _, news := range twitterNews {
		result := service.Search(news.Text, 0)

		if len(result) == 0 {
			t.Error("Twitter search error")
		}
	}
}