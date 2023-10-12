package main

import (
	"influence-hub-influencer/config"
	"influence-hub-influencer/controller"
	"influence-hub-influencer/repository"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	db := config.ConnectDb()
	repository := repository.NewRepository(db)
	controller := controller.NewController(repository)

	e := echo.New()
	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)
	e.GET("/campaign", controller.ShowCampaign)
	e.POST("/campaign", controller.ApplyCampaign)
	e.POST("/payment", controller.RequestPayment)
	e.Logger.Fatal(e.Start(":8080"))
}
