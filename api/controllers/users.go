package controllers

import (
	"app/api/config"
	"app/api/helpers"
	"app/api/modals"
	"app/request"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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
	if result.RowsAffected == 0 || !checkpass {
		return c.JSON(http.StatusBadRequest, "Kullanıcı Adı Veya Şifre Yanlış")
	}
	claims := &config.JwtCustom{
		modals.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Ballance: user.Ballance,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("mykey"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"name":     claims.User.Name,
		"email":    claims.User.Email,
		"password": claims.User.Password,
		"ballance": claims.User.Ballance,
		"token":    t,
	})
}

func UserList(c echo.Context) error {
	var user []modals.User
	db := config.Conn()
	db.Find(&user)
	return c.JSON(http.StatusOK, user)
}
