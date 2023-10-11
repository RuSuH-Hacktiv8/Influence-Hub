package main

import (
	"influence-hub-brand/config"
	"influence-hub-brand/controller"
	"influence-hub-brand/repository"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	db := config.ConnectDb()
	repository := repository.NewRepository(db)
	bc := controller.NewBrandController(repository)

	e := echo.New()
	e.POST("/register", bc.Register)
	e.POST("/login", bc.Login)
	e.Logger.Fatal(e.Start(":8081"))
}
