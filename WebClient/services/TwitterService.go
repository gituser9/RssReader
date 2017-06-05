package services

import (
	"../models"

	"log"

	"github.com/jinzhu/gorm"
)

type TwitterService struct {
	db     *gorm.DB
	config *models.Config
}

// Init - create new struct pointer with collection
func (service *TwitterService) Init(config *models.Config) *TwitterService {
	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println(err)
	}

	return &TwitterService{db: db, config: config}
}

func (service *TwitterService) GetNews(id int, page int) []models.TwitterNews {
	var result []models.TwitterNews
	offset := service.config.PageSize * (page - 1)

	service.db.Where(&models.TwitterNews{UserId: id}).
		Limit(service.config.PageSize).
		Offset(offset).
		Order("Id desc").
		Find(&result)

	return result
}

func (service *TwitterService) GetAllSources(id int) []models.TwitterSource {
	var result []models.TwitterSource

	service.db.Where(&models.TwitterSource{UserId: id}).Find(&result)

	return result
}

func (service *TwitterService) GetNewsByFilters(filters models.TwitterData) []models.TwitterNews {
	var result []models.TwitterNews
	var conditions models.TwitterNews

	if filters.SourceId != 0 {
		conditions.SourceId = filters.SourceId
	}

	service.db.Where(&conditions).Find(&result)

	return result
}
