package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserService struct {
	cfg Config
}

func CreateService(config Config) *UserService {
	service := new(UserService)
	service.cfg = config

	return service
}

func (service *UserService) GetUsersIdWithRssEnabled() []int {
	db := getDb(service.cfg)

	if db == nil {
		return nil
	}

	defer db.Close()

	var users []Users
	db.Where(&Settings{RssEnabled: true}).Select("UserId").Find(&users)

	ids := make([]int, len(users))

	for index, user := range users {
		ids[index] = user.Id
	}

	return ids
}

func getDb(config Config) *gorm.DB {
	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println("open db error:", err.Error())
		return nil
	}

	return db
}
