package services

import (
	"log"

	"github.com/jinzhu/gorm"

	"../models"
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

func (service *VkService) GetAllNews(id int) []models.VkNews {
	var result []models.VkNews

	service.db.Where(&models.VkNews{UserId: id}).Find(&result)

	return result
}

func (service *VkService) GetAllGroups(id int) []models.VkGroup {
	var result []models.VkGroup

	service.db.Where(&models.VkGroup{UserId: id}).Find(&result)

	return result
}

func (service *VkService) GetNewsByFilters(filters models.VkData) []models.VkNews {
	var result []models.VkNews
	var conditions models.VkNews

	if filters.GroupId != 0 {
		conditions.GroupId = filters.GroupId
	}

	service.db.Where(&conditions).Find(&result)

	return result
}
