package services

import (
	"crypto/aes"
	"crypto/cipher"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"fmt"

	"encoding/hex"

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

func (service *UserService) SetDb(db *gorm.DB) {
	service.db = db
}

func (service *UserService) SetConfig(cfg *models.Config) {
	service.config = cfg
}

// Auth - authorization an existing user
func (service *UserService) Auth(name, password string) *models.Users {
	var user models.Users
	service.db.Preload("Settings").Where(&models.Users{Name: name}).First(&user)
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

func (service *UserService) Update(user *models.Users) {

	// vk credentials
	/*if user.VkNewsEnabled && len(user.VkLogin) > 0 && len(user.VkPassword) > 0 {
		vkEncryptedPassword := encryptPassword(user.VkPassword)

		if len(vkEncryptedPassword) > 0 {
			user.VkPassword = vkEncryptedPassword
		}
	}*/

	service.db.Save(&user)
}

func (service *UserService) GetUser(id uint) *models.Users {
	user := models.Users{Id: id}
	service.db.Find(&user)

	return &user
}

func encryptPassword(password string) string {
	//key := []byte("AES128Key-32Characters1234567890")
	key := []byte("AES128Key-16Char")
	block, err := aes.NewCipher(key)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce, _ := hex.DecodeString("37b8e8a308c354048d245f6d")
	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	passwordBytes := []byte(password)
	encryptedPassword := aesgcm.Seal(nil, nonce, passwordBytes, nil)

	return fmt.Sprintf("%x", string(encryptedPassword))
}
