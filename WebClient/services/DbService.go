package services

import (
	"log"
	"time"

	"newshub/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var driver = ""
var connectionString = ""

func Setup(cfg models.Config) {
	driver = cfg.Driver
	connectionString = cfg.ConnectionString
}

func dbExec(closure func(db *gorm.DB)) {
	db := getDb()

	if db == nil {
		return
	}

	closure(db)
}

func getDb() *gorm.DB {
	if db != nil {
		return db
	}

	db, err := gorm.Open(driver, connectionString)

	if err != nil {
		log.Println("open db error:", err.Error())
		return nil
	}
	if driver != "sqlite3" {
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.DB().SetConnMaxLifetime(time.Hour)
	}

	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Feeds{})
	db.AutoMigrate(&models.Articles{})
	db.AutoMigrate(&models.Settings{})
	db.AutoMigrate(&models.VkNews{})
	db.AutoMigrate(&models.VkGroup{})
	db.AutoMigrate(&models.TwitterNews{})
	db.AutoMigrate(&models.TwitterSource{})

	return db
}
