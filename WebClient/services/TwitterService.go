package services

import (
	"log"
	"strconv"

	"newshub-server/models"

	"github.com/jinzhu/gorm"
)

type TwitterService struct {
	db     *gorm.DB
	config *models.Config
}

func (service *TwitterService) SetConfig(config *models.Config) {
	service.config = config
}

func (service *TwitterService) SetDb(db *gorm.DB) {
	service.db = db
}

// Init - create new struct pointer with collection
func NewTwitterService(config *models.Config) *TwitterService {
	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println(err)
	}

	return &TwitterService{db: db, config: config}
}

func (service *TwitterService) GetNews(id int64, page int, sourceId int64) []models.TwitterNewsView {
	var dbModels []models.TwitterNews
	cond := models.TwitterNews{UserId: id}
	offset := service.config.PageSize * (page - 1)

	if sourceId != 0 {
		cond.SourceId = sourceId
	}

	service.db.Where(&cond).
		Limit(service.config.PageSize).
		Offset(offset).
		Order("Id desc").
		Find(&dbModels)

	return getNewsView(dbModels)
}

func (service *TwitterService) GetAllSources(id int64) []models.TwitterSource {
	var result []models.TwitterSource

	service.db.Where(&models.TwitterSource{UserId: id}).Find(&result)

	return result
}

func (service *TwitterService) Search(searchString string, sourceId int64, userId int64) []models.TwitterNewsView {
	var dbModels []models.TwitterNews
	query := service.db.Where("Text LIKE ? and UserId = ?", "%"+searchString+"%", userId)

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
			Id:          strconv.FormatInt(item.Id, 10), // string for js
		}
	}

	return result
}
