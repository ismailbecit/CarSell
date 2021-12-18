package helpers

import (
	"app/api/config"
	"app/api/modals"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func UserInfo(c echo.Context) interface{} {
	var users modals.User
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustom)
	data := claims.User.Email
	db := config.Conn()
	db.Where("email = ?", data).Find(&users)
	email := users.Email
	return email
}
