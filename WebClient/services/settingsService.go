package services

import (
	"log"

	"github.com/jinzhu/gorm"

	"../models"
)

type SettingsService struct {
	db *gorm.DB
}

// Init - create new struct pointer with collection
func (service *SettingsService) Init(config *models.Config) *SettingsService {
	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println(err)
	}

	return &SettingsService{db: db}
}

func (service *SettingsService) Create(userId uint) {
	settings := models.Settings{UserId: userId}
	service.db.Create(&settings)
}
