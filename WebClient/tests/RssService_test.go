package tests

import (
    "testing"
    "os"
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"


    "../services"
    "../models"
)

const dbName = "test.db"
const userId = 1

var feeds models.Feeds
var user models.Users
var db *gorm.DB


func mockDb() *gorm.DB {
	db, _ = gorm.Open("sqlite3", dbName)
	db.AutoMigrate(&models.Feeds{})
	db.AutoMigrate(&models.Articles{})
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Settings{})

	feeds = models.Feeds{
		Id: 1,
		Name: "Test feed",
		Url: "test url",
		UserId: userId,
		Articles: []models.Articles {
			models.Articles{ Id: 1, Body: "Test body 1", Date: 0, FeedId: 1, IsBookmark: false, IsRead: false, Link: "Test Link 1", Title: "Test title 1" },
			models.Articles{ Id: 2, Body: "Test body 2", Date: 0, FeedId: 1, IsBookmark: true, IsRead: true, Link: "Test Link 2", Title: "Test title 2" },
		},
	}
	user = models.Users{
		Id: userId,
		Name: "Test User",
		Password: "TestPassword",
		Settings: models.Settings{
			Id: 1,
			UnreadOnly: false,
			UserId: userId,
			RssEnabled: true,
			MarkSameRead: true,
		},
	}
	db.Create(&user)
	db.Create(&feeds)

	return db
}

func fakeInit() *services.RssService {
	db := mockDb()
    settings := models.AppSettings{
        MarkSameRead:  true,
        UpdateMinutes: 30,
    }
    config := models.Config{
    	PageSize: 20,
    	UnreadOnly: false,
    	OPMLPath: ".",
	}
    service := services.RssService{}
    service.SetDb(db)
    service.SetConfig(&config)
    service.AppSettings = settings

    return &service
}

func TestGetRss(t *testing.T) {
    service := fakeInit()
    defer os.Remove(dbName)

    data := service.GetRss(userId)

	if len(data) == 0 {
		t.Error("s")
	}
}

func TestGetArticles(t *testing.T) {
	service := fakeInit()
	defer os.Remove(dbName)

	data := service.GetArticles(feeds.Id, userId, 1)

	if len(data.Articles) == 0 {
		t.Error("Articles is empty")
	}
	if len(data.Articles) != len(feeds.Articles) {
		t.Error("Articles length is wrong")
	}
}

func TestGetOneArticle(t *testing.T) {
	service := fakeInit()
	defer os.Remove(dbName)

	for _, article := range feeds.Articles {
		data := service.GetArticle(article.Id, userId)

		if data == nil {
			t.Error("Article is nil")
		}
		if data.Id != article.Id {
			t.Error("Article id is wrong")
		}
		if data.Title != article.Title {
			t.Error("Article title is wrong")
		}
	}
}

func TestDelete(t *testing.T) {
	service := fakeInit()
	defer os.Remove(dbName)

	service.Delete(feeds.Id)

	if len(service.GetRss(userId)) != 0 {
		t.Error("Feed not delete")
	}
	if len(service.GetArticles(feeds.Id, userId, 1).Articles) != 0 {
		t.Error("Articles not delete")
	}

}

func TestSetNewName(t *testing.T) {
	service := fakeInit()
	defer os.Remove(dbName)

	newName := "New Feed Name"

	service.SetNewName(newName, feeds.Id)

	if service.GetRss(userId)[0].Feed.Name != newName {
		t.Error("Name not changed")
	}
}

func TestToggleBookmark(t *testing.T) {
	service := fakeInit()
	defer os.Remove(dbName)

	for _, article := range feeds.Articles {
		isBookmark := !article.IsBookmark

		service.ToggleBookmark(article.Id, isBookmark)

		if service.GetArticle(article.Id, userId).IsBookmark != isBookmark {
			t.Error("Not bookmark toggle")
		}
	}
}

/*func TestGetBookmarks(t *testing.T) {
	service := fakeInit()
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
	service := fakeInit()
	defer os.Remove(dbName)

	service.MarkReadAll(feeds.Id)

	var articles []models.Articles
	db.Model(&models.Articles{}).Where("IsRead = 0").Find(&articles)

	if len(articles) != 0 {
		t.Error("Mark Read All error")
	}
}

// TODO: search with feed id
func TestSearch(t *testing.T) {
	service := fakeInit()
	defer os.Remove(dbName)

	for _, article := range feeds.Articles {
		result := service.Search(article.Title, article.IsBookmark, 0)

		if len(result.Articles) == 0 {
			t.Error("Not find articles for " + article.Title)
			continue
		}
		if result.Articles[0].Title != article.Title && result.Articles[0].Body != article.Body {
			t.Error("Find wrong article")
		}
	}
}

func TestToggleAsRead(t *testing.T) {
	service := fakeInit()
	defer os.Remove(dbName)

	for _, article := range feeds.Articles {
		service.ToggleAsRead(article.Id, !article.IsRead)
		var updatedArticle models.Articles
		db.Model(&models.Articles{}).Where("Id = ?", article.Id).First(&updatedArticle)

		if updatedArticle.IsRead == article.IsRead {
			t.Error("Toggle As Read error")
		}
	}
}
















/*
func TestMain(m *testing.M) {
	service = fakeInit()

	retCode := m.Run()

	os.Remove(dbName)

	os.Exit(retCode)
}*/
