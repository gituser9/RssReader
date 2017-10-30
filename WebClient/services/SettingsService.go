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
func (service SettingsService) Init(config *models.Config) *SettingsService {
	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println(err)
	}

	return &SettingsService{db: db}
}

func (service *SettingsService) SetDb(db *gorm.DB) {
	service.db = db
}

func (service *SettingsService) Create(userId uint) {
	settings := models.Settings{UserId: userId}
	service.db.Create(&settings)
}

func (service *SettingsService) Update(settings models.Settings) {
	service.db.Delete(models.Settings{UserId: settings.UserId})
	service.db.Save(&settings)
}

func (service *SettingsService) Get(userId uint) models.Settings {
	settings := models.Settings{UserId: userId}
	service.db.Find(&settings)

	return settings
}
