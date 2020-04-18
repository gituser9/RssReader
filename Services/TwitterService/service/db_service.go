package service

import (
	"log"

	"newshub-twitter-service/dao"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

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

	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println("open db error:", err.Error())
		return nil
	}

	db.AutoMigrate(&dao.TwitterNews{})
	db.AutoMigrate(&dao.TwitterSource{})
	return db
}
