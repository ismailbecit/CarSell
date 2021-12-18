package router

import (
	"app/api/config"
	"app/api/controllers"
	"app/api/controllers/ballance"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() {
	e := echo.New()
	config.Conn()

	e.POST("/register", controllers.UserRegister)
	e.POST("/login", controllers.UserLogin)

	user := e.Group("/user")
	user.Use(middleware.JWTWithConfig(config.JWTConfig))
	user.GET("/list", controllers.UserList)
	user.POST("/ballance-add", ballance.BallanceAdd)

	car := e.Group("/car")
	car.Use(middleware.JWTWithConfig(config.JWTConfig))
	car.POST("/insert", controllers.CarsInsert)
	car.POST("/delete", controllers.CarDel)
	car.POST("/buy", controllers.CarBuy)
	car.GET("/list", controllers.CarList)

	category := e.Group("/category")
	category.POST("/insert", controllers.CategoryInsert)
	category.POST("/delete", controllers.CategoryDel)
	category.GET("/list", controllers.CategoryList)
	e.Start(":8080")
}
