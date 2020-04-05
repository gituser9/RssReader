package services

import (
	"log"

	"newshub/models"

	"github.com/jinzhu/gorm"
)

// VkService - service
type VkService struct {
	db     *gorm.DB
	config *models.Config
}

// Init - create new struct pointer with collection
func NewVkService(config *models.Config) *VkService {
	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println(err)
	}

	return &VkService{db: db, config: config}
}

func (service *VkService) SetDb(db *gorm.DB) {
	service.db = db
}

func (service *VkService) SetConfig(cfg *models.Config) {
	service.config = cfg
}

func (service *VkService) GetNews(id int64, page int, groupId int64) []models.VkNews {
	var result []models.VkNews
	conditions := models.VkNews{
		UserId: id,
	}
	offset := service.config.PageSize * (page - 1)

	if groupId != 0 {
		conditions.GroupId = groupId
	}

	service.db.Where(&conditions).
		Limit(service.config.PageSize).
		Offset(offset).
		Order("Id desc").
		Find(&result)

	return result
}

func (service *VkService) GetAllGroups(id int64) []models.VkGroup {
	var result []models.VkGroup

	service.db.Where(&models.VkGroup{UserId: id}).Find(&result)

	return result
}

func (service *VkService) Search(searchString string, groupId int64, userId int64) []models.VkNews {
	var result []models.VkNews
	query := service.db.Where("Text LIKE ? and userId = ?", "%"+searchString+"%", userId)

	if groupId != 0 {
		query = query.Where(&models.VkNews{GroupId: groupId})
	}

	query.Order("Id desc").Find(&result)

	return result
}
