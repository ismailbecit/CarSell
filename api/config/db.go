package config

import (
	"app/api/modals"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Conn() *gorm.DB {
	dsn := "root:aea76026@tcp(127.0.0.1:3306)/carsell?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Veri Tabanına Bağlanılamadı")
	}
	db.AutoMigrate(
		&modals.User{},
		&modals.Category{},
		&modals.Car{},
	)
	return db
}
