package router

import (
	"app/api/config"
	"app/api/controllers"

	"github.com/labstack/echo/v4"
)

func Router() {
	e := echo.New()
	config.Conn()

	user := e.Group("/user")
	user.POST("/register", controllers.UserRegister)
	user.POST("/login", controllers.UserLogin)
	user.GET("/list", controllers.UserList)

	car := e.Group("/car")
	car.POST("/insert", controllers.CarsInsert)
	car.POST("/delete", controllers.CarDel)
	car.GET("/list", controllers.CarList)

	category := e.Group("/category")
	category.POST("/insert", controllers.CategoryInsert)
	category.POST("/delete", controllers.CategoryDel)
	category.GET("/list", controllers.CategoryList)
	e.Start(":8080")
}
