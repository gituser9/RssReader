package services

import (
	"log"

	"newshub/models"

	"github.com/jinzhu/gorm"
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

func (service *SettingsService) Create(userId int64) {
	settings := models.Settings{UserId: userId}
	service.db.Create(&settings)
}

func (service *SettingsService) Update(settings models.Settings) {
	service.db.Delete(models.Settings{UserId: settings.UserId})
	service.db.Save(&settings)
}

func (service *SettingsService) Get(userId int64) models.Settings {
	settings := models.Settings{UserId: userId}
	service.db.Find(&settings)

	return settings
}
