package services

import (
	"log"

	"../models"
	"github.com/jinzhu/gorm"
)

// VkService - service
type VkService struct {
	db     *gorm.DB
	config *models.Config
}

// Init - create new struct pointer with collection
func (service *VkService) Init(config *models.Config) *VkService {
	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println(err)
	}

	return &VkService{db: db, config: config}
}

func (service *VkService) GetAll(id int) []models.VkNews {
	var result []models.VkNews

	service.db.Where(models.VkNews{UserId: id}).Find(&result)

	return result
}
