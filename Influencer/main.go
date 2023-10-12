package main

import (
	"influence-hub-influencer/config"
	"influence-hub-influencer/controller"
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

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("00:00").Do(controller.UpdateFollowerCount)
	s.StartAsync()

	e := echo.New()
	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)
	e.Logger.Fatal(e.Start(":8080"))
}
