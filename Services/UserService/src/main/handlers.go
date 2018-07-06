package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type HttpHandler struct {
	userService *UserService
}

func CreateHttphandler(config Config) *HttpHandler {
	handler := new(HttpHandler)
	handler.userService = CreateService(config)

	return handler
}

func (handler HttpHandler) GetUsersIdWithRssEnabled(ctx *fasthttp.RequestCtx) {
	ids := handler.userService.GetUsersIdWithRssEnabled()
	json.NewEncoder(ctx).Encode(ids)
}
