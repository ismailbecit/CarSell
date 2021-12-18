package middleware

import (
	"app/api/config"
	"app/api/modals"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) modals.User {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustom)
	info := modals.User{
		Name:     claims.User.Name,
		Email:    claims.User.Email,
		Password: claims.User.Password,
		Ballance: claims.User.Ballance,
	}
	return info
}
