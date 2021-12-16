package controllers

import (
	"app/api/config"
	"app/api/modals"
	"app/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CategoryInsert(c echo.Context) error {
	var rq request.CategoryInsert
	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := config.Conn()
	db.Create(&modals.Category{
		Year:  rq.Year,
		Price: rq.Price,
		Name:  rq.Name,
	})
	return c.JSON(http.StatusOK, "Tebrikler Kategori Başarıyla Oluşturuldu")

}
func CategoryDel(c echo.Context) error {
	var rq request.CategoryDel
	var category modals.Category

	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := config.Conn()
	result := db.Find(&category, rq.ID)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "Böyle Bir Kayıt Bulunamadı")
	}
	db.Delete(&category, rq.ID)
	return c.JSON(http.StatusBadRequest, "Kayıt Başarıyla Silindi")
}
func CategoryList(c echo.Context) error {
	var category []modals.Category
	db := config.Conn()
	db.Find(&category)
	return c.JSON(http.StatusOK, category)
}
