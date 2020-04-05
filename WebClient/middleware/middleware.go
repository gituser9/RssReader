package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"newshub/models"

	"github.com/dgrijalva/jwt-go"
)

const authHeader = "Authorization"
const bearer = "Bearer "

type AuthenticationMiddleware struct {
	allowRoutes map[string]bool // todo: config
	config      models.Config
}

// Initialize it somewhere
func (amw *AuthenticationMiddleware) Populate(cfg models.Config) {
	amw.config = cfg
	amw.allowRoutes = map[string]bool{
		"/":              true,
		"/auth":          true,
		"/registration":  true,
		"/users/refresh": true,
	}
}

// Middleware function, which will be called for each request
func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.String()

		if _, ok := amw.allowRoutes[path]; ok {
			next.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(path, "/dist") {
			next.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(path, "/static") {
			next.ServeHTTP(w, r)
			return
		}
		if err := amw.jwtValidate(r); err != nil {
			log.Println("jwt validation error:", err)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (amw *AuthenticationMiddleware) jwtValidate(r *http.Request) error {
	tokenString := getJwtString(r)

	if tokenString == "" {
		return errors.New("token is empty")
	}

	claims := models.JwtClaims{}
	_, err := jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(amw.config.JwtSign), nil
		},
	)

	if err != nil {
		log.Println("JWT error:", err)
		return fmt.Errorf("parse JWT error: %s", err)
	}
	if claims.Exp == 0 || claims.Exp < time.Now().Unix() {
		return errors.New("JWT is expired")
	}

	return nil
}

func getJwtString(r *http.Request) string {
	return strings.TrimSpace(strings.TrimPrefix(r.Header.Get(authHeader), bearer))
}
