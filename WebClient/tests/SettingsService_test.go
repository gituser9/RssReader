package tests

import (
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"../models"
	"../services"
)

func makeDbForSettings(withUser bool) *gorm.DB {
	db, _ = gorm.Open("sqlite3", dbName)
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Settings{})

	user = models.Users{
		Id:       userId,
		Name:     "Test User",
		Password: "TestPassword",
	}

	if withUser {
		user.Settings = models.Settings{
			Id:     1,
			UserId: userId,
		}
	}

	db.Create(&user)

	return db
}

func fakeInitSettingsService(withUser bool) *services.SettingsService {
	db := makeDbForSettings(withUser)
	service := services.SettingsService{}
	service.SetDb(db)

	return &service
}

func TestCreate(t *testing.T) {
	log.Println("Start TestCreate")
	service := fakeInitSettingsService(false)
	defer os.Remove(dbName)

	var settings models.Settings
	service.Create(user.Id)
	db.Where(&models.Settings{UserId: user.Id}).Find(&settings)

	if settings.Id == 0 {
		t.Error("Create settings error")
	}
}

func TestSettingsUpdate(t *testing.T) {
	log.Println("Start TestSettingsUpdate")
	service := fakeInitSettingsService(true)
	defer os.Remove(dbName)

	var result models.Settings
	newSettings := models.Settings{
		RssEnabled: true,
		UserId:     user.Id,
	}
	service.Update(newSettings)
	db.Where(&models.Settings{UserId: user.Id}).Find(&result)

	if !result.RssEnabled {
		t.Error("Update error")
	}
}

/*func TestGetSettings(t *testing.T) {
	log.Println("Start TestGetSettings")
	service := fakeInitSettingsService(true)
	defer os.Remove(dbName)

	result := service.Get(user.Id)

	if result.Id == 0 {
		t.Error("Get Error")
	}
}*/
