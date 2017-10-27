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

var vkNews = make([]models.VkNews, 10)

const vkUserId int = 1

func mockDbForVk() *gorm.DB {
	db, _ = gorm.Open("sqlite3", dbName)
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Settings{})
	db.AutoMigrate(&models.VkGroup{})
	db.AutoMigrate(&models.VkNews{})

	user = models.Users{
		//Id: userId,
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
		vkGroup := models.VkGroup{
			Gid:        i,
			UserId:     vkUserId,
			Name:       "VkGroup " + strconv.Itoa(i),
			LinkedName: "vkgroup_" + strconv.Itoa(i),
			Image:      "image",
		}
		db.Create(&vkGroup)

		for j := 1; j < 11; j++ {
			stringNum := strconv.Itoa(j + (i * 10))
			vkItem := models.VkNews{
				UserId:    vkUserId,
				GroupId:   vkGroup.Id,
				PostId:    j,
				Text:      "VkNews Text " + stringNum,
				Image:     "Image " + stringNum,
				Link:      "Link " + stringNum,
				Timestamp: 0,
			}
			db.Create(&vkItem)
		}
		db.Find(&vkNews)
	}

	return db
}

func fakeInitVkService() *services.VkService {
	db := mockDbForVk()
	config := models.Config{
		PageSize:   20,
		UnreadOnly: false,
		OPMLPath:   ".",
	}
	service := services.VkService{}
	service.SetDb(db)
	service.SetConfig(&config)

	return &service
}

func TestGetVkNews(t *testing.T) {
	log.Println("Start TestGetVkNews")
	service := fakeInitVkService()
	defer os.Remove(dbName)

	news := service.GetNews(vkUserId, 1)

	if len(news) == 0 {
		t.Error("News is empty")
	}
}

// TODO: different users
func TestGetAllVkGroups(t *testing.T) {
	log.Println("Start TestGetAllVkGroups")
	service := fakeInitVkService()
	defer os.Remove(dbName)

	groups := service.GetAllGroups(vkUserId)

	if len(groups) == 0 {
		t.Error("Empty groups")
	}
}

func TestGetVkByFilters(t *testing.T) {
	log.Println("Start TestGetVkByFilters")
	service := fakeInitVkService()
	defer os.Remove(dbName)

	for _, item := range vkNews {
		filters := models.VkData{
			GroupId: item.GroupId,
		}
		news := service.GetNewsByFilters(filters)

		if len(news) == 0 {
			t.Error("News is empty")
			continue
		}

		for _, newsItem := range news {
			if newsItem.GroupId != filters.GroupId {
				t.Error("Filter by group error")
			}
		}

		filters.SearchString = item.Text
		filters.GroupId = 0
		news = service.GetNewsByFilters(filters)

		if len(news) == 0 {
			t.Error("Filter by string error")
		}
	}
}

// TODO: groups
func TestVkSearch(t *testing.T) {
	log.Println("Start TestVkSearch")
	service := fakeInitVkService()
	defer os.Remove(dbName)

	for _, news := range vkNews {
		result := service.Search(news.Text, 0)

		if len(result) == 0 {
			t.Error("Vk search error")
		}
	}
}
