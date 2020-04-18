package service

import "newshub-twitter-service/dao"

var config = dao.Config{}
var updChan chan dao.User
var closeChan chan bool

func Setup(cfg dao.Config) {
	config = cfg

	updChan = make(chan dao.User, 4)
	closeChan = make(chan bool, 1)
	go listener()
}
