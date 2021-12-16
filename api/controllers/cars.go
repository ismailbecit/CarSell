package controllers

import (
	"app/api/config"
	"app/api/modals"
	"app/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CarList(c echo.Context) error {
	var cars []modals.Car
	db := config.Conn()
	db.Preload("User").Preload("Category").Find(&cars)
	return c.JSON(http.StatusOK, cars)
}

func CarsInsert(c echo.Context) error {
	var rq request.CarInsert

	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// kayıt işlemi
	db := config.Conn()
	db.Create(&modals.Car{
		Userfk:     rq.Userfk,
		Categoryfk: rq.Categoryfk,
	})
	return c.JSON(http.StatusOK, "Kayıt İşlenmi Başarılı")
}
func CarDel(c echo.Context) error {
	var rq request.CarDel
	var car modals.Car
	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := config.Conn()
	// boyle bir kayıt varmı sorgulama
	result := db.First(&car, rq.ID)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "Böyle Bir Kayıt Bulunamadı")
	}

	// silme islemi
	db.Delete(&car, rq.ID)
	return c.JSON(http.StatusOK, "Silme İşlemi Başarıyla Gerçeşleştirildi")
}
