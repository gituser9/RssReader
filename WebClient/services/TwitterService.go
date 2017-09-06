package services

import (
	"../models"

	"log"

	"strconv"

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

func (service *TwitterService) GetNews(id int, page int) []models.TwitterNewsView {
	var dbModels []models.TwitterNews
	offset := service.config.PageSize * (page - 1)

	service.db.Where(&models.TwitterNews{UserId: id}).
		Limit(service.config.PageSize).
		Offset(offset).
		Order("Id desc").
		Find(&dbModels)

	return getNewsView(dbModels)
}

func (service *TwitterService) GetAllSources(id int) []models.TwitterSource {
	var result []models.TwitterSource

	service.db.Where(&models.TwitterSource{UserId: id}).Find(&result)

	return result
}

func (service *TwitterService) GetNewsByFilters(filters models.TwitterData) []models.TwitterNewsView {
	var dbModels []models.TwitterNews
	var conditions models.TwitterNews

	if filters.SourceId != 0 {
		conditions.SourceId = filters.SourceId
	}

	service.db.Where(&conditions).Order("Id desc").Find(&dbModels)

	return getNewsView(dbModels)
}

func (service *TwitterService) Search(searchString string, sourceId int) []models.TwitterNewsView {
	var dbModels []models.TwitterNews
	query := service.db.Where("Text LIKE ?", "%"+searchString+"%")

	if sourceId != 0 {
		query = query.Where(&models.TwitterNews{SourceId: sourceId})
	}

	query.Order("Id desc").Find(&dbModels)

	return getNewsView(dbModels)
}

func getNewsView(dbModels []models.TwitterNews) []models.TwitterNewsView {
	result := make([]models.TwitterNewsView, len(dbModels))

	for index, item := range dbModels {
		result[index] = models.TwitterNewsView{
			SourceId:    item.SourceId,
			ExpandedUrl: item.ExpandedUrl,
			Image:       item.Image,
			Text:        item.Text,
			Id:          strconv.FormatUint(item.Id, 10),
		}
	}

	return result
}
