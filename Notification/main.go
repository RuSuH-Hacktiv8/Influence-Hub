package main

import (
	"influence-hub-notification/handler"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	m := e.Group("/mails")
	m.POST("/success_register", handler.SuccessRegister)
	m.POST("/success_apply_campaign", handler.SuccessApplyCampaign)
	m.POST("/payment_campaign_received", handler.PaymentReceived)

	e.Logger.Fatal(e.Start(":8082"))
}
