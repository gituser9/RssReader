package controllers

import (
	"log"
	"net/http"
	"strings"

	"newshub/models"

	"github.com/dgrijalva/jwt-go"
)

const authHeader = "Authorization"
const bearer = "Bearer "

var Config *models.Config

func getClaims(r *http.Request) models.JwtClaims {
	claims := models.JwtClaims{}
	token := getJwtString(r)

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.JwtSign), nil
	})

	if err != nil {
		log.Println("parse JWT error:", err)
	}

	return claims
}

func getJwtString(r *http.Request) string {
	return strings.TrimSpace(strings.TrimPrefix(r.Header.Get(authHeader), bearer))
}

func getInclude(include string) []string {
	return strings.Split(include, ",")
}
