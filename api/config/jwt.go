package config

import (
	"app/api/modals"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustom struct {
	User modals.User `json:"user"`
	jwt.StandardClaims
}

type JwtCarCustom struct {
	Car modals.Car `json:"car"`
	jwt.StandardClaims
}

var JWTConfig = middleware.JWTConfig{
	Claims:     &JwtCustom{},
	SigningKey: []byte("mykey"),
}
