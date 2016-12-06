package services

import (
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"../models"
)

// RssService - service
type UserService struct {
	db     *gorm.DB
	config *models.Config
}

// Init - create new struct pointer with collection
func (service *UserService) Init(config *models.Config) *UserService {
	db, err := gorm.Open(config.Driver, config.ConnectionString)

	if err != nil {
		log.Println(err)
	}

	return &UserService{db: db, config: config}
}

// Auth - authorization an existing user
func (service *UserService) Auth(name, password string) *models.Users {
	var user models.Users
	service.db.Where(&models.Users{Name: name}).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil
	}

	user.Password = ""
	return &user
}

// Register - create new user
func (service *UserService) Register(name, password string) models.RegistrationData {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("Registration error: ", err.Error())
		return models.RegistrationData{User: nil, Message: "Create password error"}
	}

	user := models.Users{Name: name}
	service.db.Where(&user).Find(&user)

	if user.Id != 0 {
		return models.RegistrationData{User: nil, Message: "User with this name already exist"}
	}

	settingsService := new(SettingsService).Init(service.config)
	user.Password = string(hashedPassword)
	service.db.Create(&user)
	service.db.Find(&user) // for get id
	settingsService.Create(user.Id)
	user.Password = ""

	return models.RegistrationData{User: &user, Message: ""}
}
