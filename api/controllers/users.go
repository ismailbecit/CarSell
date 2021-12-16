package controllers

import (
	"app/api/config"
	"app/api/helpers"
	"app/api/modals"
	"app/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserRegister(c echo.Context) error {
	var user modals.User
	var rq request.UserInsert
	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// veri tabanında aynı mailde bir kayıt var mı ?
	db := config.Conn()
	row := db.Where("email = ? ", rq.Email).Find(&user)
	if row.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "Hata! Böyle Bir Kullanıcı Zaten Kayıtlı")
	}
	if rq.Email == "" || rq.Name == "" || rq.Password == "" {
		return c.JSON(http.StatusBadRequest, "Hata! Tüm Verileri Eksiksiz Doldurun")
	}
	hashpass, _ := helpers.HashPassword(rq.Password)
	// veri tabanına kayıt ettirelim
	db.Create(&modals.User{
		Name:     rq.Name,
		Password: hashpass,
		Email:    rq.Email,
	})
	return c.JSON(http.StatusOK, "Kayıt Başarıyla Tamamlandı")
}

func UserLogin(c echo.Context) error {
	var user modals.User
	var rq request.UserLogin
	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if rq.Email == "" || rq.Password == "" {
		return c.JSON(http.StatusBadRequest, "Lütfen Tüm Verileri Eksiksiz Giriniz!")
	}
	db := config.Conn()

	// verilerle eşleşen kayıt var mı ?
	result := db.Where("email = ?", rq.Email).Find(&user)
	checkpass := helpers.CheckPasswordHash(rq.Password, user.Password)
	if result.RowsAffected == 0 || checkpass == false {
		return c.JSON(http.StatusBadRequest, "Kullanıcı Adı Veya Şifre Yanlış")
	}
	return c.JSON(http.StatusOK, user)
}

func UserList(c echo.Context) error {
	var user []modals.User
	db := config.Conn()
	db.Find(&user)
	return c.JSON(http.StatusOK, user)
}
