package main

import (
	"influence-hub-brand/config"
	"influence-hub-brand/controller"
	"influence-hub-brand/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.ConnectDb()
	repository := repository.NewRepository(db)
	controller := controller.NewController(repository)

	e := echo.New()
	e.POST("/register", controller.Register)
	e.Logger.Fatal(e.Start(":8081"))
}
