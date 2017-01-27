package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
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
func (service *UserService) saveVkCredentials(id uint, login string, password string) bool {
	// create hash
	encryptedPassword := encryptPassword(password)

	if len(encryptedPassword) == 0 {
		log.Println("Encrypt Vk password error")
		return false
	}

	// get user
	user := models.Users{Id: id}
	service.db.Find(&user)

	// update user
	user.VkLogin = login
	user.VkPassword = encryptedPassword
	user.VkNewsEnabled = true

	service.db.Save(&user)

	return true
}

func encryptPassword(password string) string {
	key := []byte("AES256Key-32Characters1234567890")

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println(err.Error())
		return ""
	}

	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	passwordBytes := []byte(password)
	encryptedPassword := aesgcm.Seal(nil, nonce, passwordBytes, nil)
	log.Printf("%x\n", encryptedPassword)

	return string(encryptedPassword)
}
