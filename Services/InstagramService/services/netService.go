package service

import (
	"encoding/json"
	"net/http"
)

type NetService interface {
}

func (service *NetService) getToken() string {
	return ""
}

func (service *NetService) getNews() {

}
