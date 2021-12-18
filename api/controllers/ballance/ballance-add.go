package ballance

import (
	"app/api/config"
	"app/api/helpers"
	"app/api/modals"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func BallanceAdd(c echo.Context) error {
	var user modals.User
	amount, _ := strconv.Atoi(c.QueryParam("amount"))
	email := helpers.UserInfo(c)
	db := config.Conn()
	db.Where("email = ? ", email).First(&user)

	amount += int(user.Ballance)
	db.Model(&user).Where("email = ? ", email).Update("ballance", amount)

	return c.JSON(200, fmt.Sprintf("Bakiye Hesabınıza Eklendi Toplam Bakiyeniz: %v", amount))

}
