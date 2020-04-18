package services

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"newshub-server/models"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// RssService - service
type UserService struct {
	config *models.Config
}

func NewUserService(cfg *models.Config) *UserService {
	return &UserService{config: cfg}
}

func (service *UserService) SetConfig(cfg *models.Config) {
	service.config = cfg
}

// Auth - authorization an existing user
func (service *UserService) Auth(name, password string) *models.Users {
	user := new(models.Users)

	dbExec(func(db *gorm.DB) {
		if err := db.Where(models.Users{Name: name}).First(&user).Error; err != nil {
			log.Println("get user error:", err)
		}
	})

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
		return user
	}

	return nil
}

// Register - create new user
func (service *UserService) Register(name, password string) (models.Users, error) {
	existingUser := models.Users{}
	user := models.Users{}
	var userError error = nil

	dbExec(func(db *gorm.DB) {
		if err := db.Where(models.Users{Name: name}).First(&existingUser); err == nil {
			userError = errors.New("Логин уже занят")
			return
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &models.Users{
			Name:     name,
			Password: string(hashedPassword),
		}
		if err := db.Save(&user).Error; err != nil {
			log.Println("create user error:", err)
			userError = errors.New("Произошла неизвестная ошибка")
			return
		}

		settings := models.Settings{
			UserId:       user.Id,
			MarkSameRead: true,
		}
		db.Save(&settings)
	})

	return user, userError
}

func (service *UserService) Update(user *models.Users) {

	// vk credentials
	/*if user.VkNewsEnabled && len(user.VkLogin) > 0 && len(user.VkPassword) > 0 {
		vkEncryptedPassword := encryptPassword(user.VkPassword)

		if len(vkEncryptedPassword) > 0 {
			user.VkPassword = vkEncryptedPassword
		}
	}*/

	dbExec(func(db *gorm.DB) {
		if err := db.Save(&user).Error; err != nil {
			log.Println("update user error:", err)
		}
	})

}

func (service *UserService) GetUser(id int64) models.Users {
	user := models.Users{Id: id}

	dbExec(func(db *gorm.DB) {
		db.Find(&user)
	})

	return user
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
