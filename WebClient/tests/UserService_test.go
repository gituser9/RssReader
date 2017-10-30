package tests

import (
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"

	"../models"
	"../services"
)

const userPassword = "TestPassword"

func makeBdForUser(noUser bool) *gorm.DB {
	db, _ = gorm.Open("sqlite3", dbName)
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Settings{})

	if noUser {
		return db
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)

	user = models.Users{
		Id:       userId,
		Name:     "Test User",
		Password: string(hashedPassword),
		Settings: models.Settings{
			Id:           1,
			UnreadOnly:   false,
			UserId:       userId,
			RssEnabled:   true,
			MarkSameRead: true,
		},
	}
	db.Create(&user)

	return db
}

func fakeInitUserService(noUser bool) *services.UserService {
	db := makeBdForUser(noUser)
	config := models.Config{
		PageSize:   20,
		UnreadOnly: false,
		OPMLPath:   ".",
	}
	service := services.UserService{}
	service.SetDb(db)
	service.SetConfig(&config)

	return &service
}

func TestAuth(t *testing.T) {
	log.Println("Start TestAuth")
	service := fakeInitUserService(false)
	defer os.Remove(dbName)

	result := service.Auth(user.Name, userPassword)

	if result == nil {
		t.Error("Auth error")
	}
}

/*func TestRegister(t *testing.T) {
	log.Println("Start TestRegister")
	service := fakeInitUserService(true)
	defer os.Remove(dbName)

	var settings models.Settings
	name := "TestName"
	result := service.Register(name, userPassword)
	db.Find(&settings)

	if result.User == nil || len(result.Message) > 0 {
		t.Error("Registration error")
	}
	if result.User.Name != name {
		t.Error("Wrong name after registration")
	}
	if settings.UserId == 0 {
		t.Error("Settings is not created")
	}
}*/

func TestUpdate(t *testing.T) {
	log.Println("Start TestUpdate")
	service := fakeInitUserService(false)
	defer os.Remove(dbName)

	newName := "NewTestName"
	newUserData := models.Users{
		Id:   user.Id,
		Name: newName,
	}
	service.Update(&newUserData)
	db.Find(&user)

	if user.Name != newName {
		t.Error("Update error")
	}
}

func TestGetUser(t *testing.T) {
	log.Println("Start TestGetUser")
	service := fakeInitUserService(false)
	defer os.Remove(dbName)

	result := service.GetUser(user.Id)

	if result.Id != user.Id || result.Name != user.Name {
		t.Error("Get user error")
	}
}
