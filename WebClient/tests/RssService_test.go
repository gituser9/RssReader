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

const dbName = "test.db"
const userId uint = 1

var feeds []models.Feeds = make([]models.Feeds, 10)
var user models.Users
var db *gorm.DB

func mockDbForRss() *gorm.DB {
	db, _ = gorm.Open("sqlite3", dbName)
	db.AutoMigrate(&models.Feeds{})
	db.AutoMigrate(&models.Articles{})
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Settings{})

	user = models.Users{
		Id:       userId,
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
		feedIdString := strconv.Itoa(i)
		feed := models.Feeds{
			Name:   "Test feed " + feedIdString,
			Url:    "test url " + feedIdString,
			UserId: userId,
		}
		db.Create(&feed)

		for j := 1; j < 11; j++ {
			stringNum := strconv.Itoa(j + (i * 10))
			article := models.Articles{Body: "Test body ", Date: 0, FeedId: feed.Id, IsBookmark: j%2 == 0, IsRead: j%2 == 0, Link: "Test Link " + stringNum, Title: "Test title " + stringNum}

			db.Create(&article)
		}

		db.Preload("Articles").Find(&feeds)
	}

	return db
}

func fakeInitRssService() *services.RssService {
	db := mockDbForRss()
	settings := models.AppSettings{
		MarkSameRead:  true,
		UpdateMinutes: 30,
	}
	config := models.Config{
		PageSize:   20,
		UnreadOnly: false,
		OPMLPath:   ".",
	}
	service := services.RssService{}
	service.SetDb(db)
	service.SetConfig(&config)
	service.AppSettings = settings

	return &service
}

func TestGetRss(t *testing.T) {
	log.Println("start TestGetRss")
	service := fakeInitRssService()
	defer os.Remove(dbName)

	data := service.GetRss(userId)

	if len(data) == 0 {
		t.Error("Get rss error")
	}
}

func TestGetArticles(t *testing.T) {
	log.Println("start TestGetArticles")
	service := fakeInitRssService()
	defer os.Remove(dbName)

	for _, feed := range feeds {
		data := service.GetArticles(feed.Id, userId, 1)

		if len(data.Articles) == 0 {
			t.Error("Articles is empty for: " + feed.Name)
		}
		if len(data.Articles) != len(feed.Articles) {
			t.Error("Articles length is wrong for: " + feed.Name)
		}
	}
}

func TestGetOneArticle(t *testing.T) {
	log.Println("start TestGetOneArticle")
	service := fakeInitRssService()
	defer os.Remove(dbName)

	for _, feed := range feeds {
		for _, article := range feed.Articles {
			data := service.GetArticle(article.Id, userId)

			if data == nil {
				t.Error("Article is nil for: " + feed.Name + " - " + article.Title)
			}
			if data.Id != article.Id {
				t.Error("Article id is wrong for: " + feed.Name + " - " + article.Title)
			}
			if data.Title != article.Title {
				t.Error("Article title is wrong for: " + feed.Name + " - " + article.Title)
			}
		}
	}
}

func TestDelete(t *testing.T) {
	log.Println("start TestDelete")
	service := fakeInitRssService()
	defer os.Remove(dbName)

	for _, feed := range feeds {
		service.Delete(feed.Id)
		var testFeeds models.Feeds
		db.Where(&models.Feeds{Id: feed.Id}).Find(&testFeeds)

		if testFeeds.Id != 0 {
			t.Error("Feed not delete: " + feed.Name)
		}
		if len(service.GetArticles(feed.Id, userId, 1).Articles) != 0 {
			t.Error("Articles not delete for: " + feed.Name)
		}
	}
}

func TestSetNewName(t *testing.T) {
	log.Println("start TestSetNewName")
	service := fakeInitRssService()
	defer os.Remove(dbName)

	for index, feed := range feeds {
		newName := "New Feed Name " + strconv.Itoa(index)

		service.SetNewName(newName, feed.Id)
		var testFeeds models.Feeds
		db.Where(&models.Feeds{Id: feed.Id}).Find(&testFeeds)

		if testFeeds.Name != newName {
			t.Error("Name not changed for: " + feed.Name)
		}
	}
}

func TestToggleBookmark(t *testing.T) {
	log.Println("start TestToggleBookmark")
	service := fakeInitRssService()
	defer os.Remove(dbName)

	for _, feed := range feeds {
		for _, article := range feed.Articles {
			isBookmark := !article.IsBookmark

			service.ToggleBookmark(article.Id, isBookmark)

			if service.GetArticle(article.Id, userId).IsBookmark != isBookmark {
				t.Error("Not bookmark toggle for: " + feed.Name + " - " + article.Title)
			}
		}
	}
}

/*func TestGetBookmarks(t *testing.T) {
log.Println("start )
	service := fakeInitRssService()
	defer os.Remove(dbName)

	bookmarkCount := 0

	for _, article := range feeds.Articles {
		if article.IsBookmark {
			bookmarkCount++
		}
	}

	bookmarks := service.GetBookmarks(1, userId)

	if len(bookmarks.Articles) != bookmarkCount {
		t.Error("Wrong bookmarks count")
	}
}*/

func TestMarkReadAll(t *testing.T) {
	log.Println("start TestMarkReadAll")
	service := fakeInitRssService()
	defer os.Remove(dbName)

	for _, feed := range feeds {
		service.MarkReadAll(feed.Id)

		var articles []models.Articles
		db.Model(&models.Articles{}).Where("FeedId = ? and IsRead = 0", feed.Id).Find(&articles)

		if len(articles) != 0 {
			t.Error("Mark Read All error for: " + feed.Name)
		}
	}
}

// TODO: search with feed id
func TestSearch(t *testing.T) {
	log.Println("start TestSearch")
	service := fakeInitRssService()
	defer os.Remove(dbName)

	for _, feed := range feeds {
		for _, article := range feed.Articles {
			result := service.Search(article.Title, article.IsBookmark, 0)

			if len(result.Articles) == 0 {
				t.Error("Not find articles for: " + feed.Name + " - " + article.Title)
				continue
			}
			if result.Articles[0].Title != article.Title && result.Articles[0].Body != article.Body {
				t.Error("Find wrong article for: " + feed.Name + " - " + article.Title)
			}
		}
	}
}

func TestToggleAsRead(t *testing.T) {
	log.Println("start TestToggleAsRead")
	service := fakeInitRssService()
	defer os.Remove(dbName)

	for _, feed := range feeds {
		for _, article := range feed.Articles {
			service.ToggleAsRead(article.Id, !article.IsRead)
			var updatedArticle models.Articles
			db.Model(&models.Articles{}).Where("Id = ?", article.Id).First(&updatedArticle)

			if updatedArticle.IsRead == article.IsRead {
				t.Error("Toggle As Read error for: " + feed.Name + " - " + article.Title)
			}
		}
	}
}

/*
func TestMain(m *testing.M) {
	service = fakeInitRssService()

	retCode := m.Run()

	os.Remove(dbName)

	os.Exit(retCode)
}*/
