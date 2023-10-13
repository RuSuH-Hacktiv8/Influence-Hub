package main

import (
	"influence-hub-influencer/config"
	"influence-hub-influencer/controller"
	"influence-hub-influencer/middleware"
	"influence-hub-influencer/repository"
	"time"

	"github.com/go-co-op/gocron"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	db := config.ConnectDb()
	repository := repository.NewRepository(db)
	controller := controller.NewController(repository)
	auth := middleware.NewAuth(*repository)
  
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("00:00").Do(controller.UpdateFollowerCount)
	s.StartAsync()

	e := echo.New()
	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)
	e.GET("/campaign", controller.ShowCampaign, auth.AuthUser)
	e.POST("/campaign", controller.ApplyCampaign, auth.AuthUser)
	e.POST("/payment", controller.RequestPayment, auth.AuthUser)
	e.Logger.Fatal(e.Start(":8080"))
}
