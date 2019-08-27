package main

import (
	"my-trophy-prototype-api/config"
	"my-trophy-prototype-api/handler"

	"my-trophy-prototype-api/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	// load config
	c := config.NewConfig()

	// create Echo
	e := echo.New()

	// Set middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS(c.AllowOrigins))

	// Create database connection
	db, err := gorm.Open("mysql", c.DBUserName+":"+c.DBPassword+"@tcp("+c.DBHost+":"+c.DBPort+")/mytrophy?charset=utf8mb4&collation=utf8mb4_bin&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize handler
	h := &handler.Handler{DB: db}

	// Bind routes
	e.GET("/users/", h.FindUsers)
	e.GET("/users/:id", h.FindUserByID)
	e.GET("/trophies/", h.FindTrophies)
	e.GET("/trophies/:id", h.FindTrophyByID)

	// Start server
	e.Logger.Fatal(e.Start(":" + c.PORT))
}
