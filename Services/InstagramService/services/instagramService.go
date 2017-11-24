package service

import (
    "../models"
)

type InstagramService struct {
    cfg *models.Configuration
}

func (service *InstagramService) Update(data string) {

}

func InitInstagramService(cfg *models.Configuration) *InstagramService {
    instagramservice := new(InstagramService)
    instagramservice.cfg = cfg

    return instagramservice
}