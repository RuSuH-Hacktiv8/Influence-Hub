package controller

import (
	"influence-hub-influencer/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cn *Controller) Login(c echo.Context) error {
	influencer := new(models.Influencer)
	if err := c.Bind(influencer); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
}
