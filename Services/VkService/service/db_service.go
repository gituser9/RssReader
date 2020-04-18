package service

import (
	"log"
	"newshub-vk-service/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var driver = ""
var connectionString = ""

func Setup(cfg model.Config) {
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

	return db
}
