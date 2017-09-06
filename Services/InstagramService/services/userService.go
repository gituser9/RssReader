package service

import (
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "../models"
)

type UserService struct {
    cfg *models.Configuration
}

func (service *UserService) GetUsers() []models.Users {
    var users []models.Users
    db, err := gorm.Open(service.cfg.Driver, service.cfg.ConnectionString)

    if err != nil {
        log.Println("Database open" + err.Error())
        return users
    }

    var settings []models.Settings
    db.Where(&models.Settings{InstagramEnabled:true}).Select("UserId").Find(&settings)

    ids := make([]int, len(settings))

    for i := 0; i < len(settings); i++  {
        ids[i] = settings[i].UserId
    }

    db.Where(ids).Find(&users)

    //log.Println(users)

    return users
}

func InitUserService(cfg *models.Configuration) *UserService {
    userService := new(UserService)
    userService.cfg = cfg

    return userService
}
