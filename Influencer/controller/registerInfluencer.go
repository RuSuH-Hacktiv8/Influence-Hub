package controller

import (
	"influence-hub-influencer/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (cn *Controller) Register(c echo.Context) error {
	influencer := new(models.Influencer)
	if err := c.Bind(influencer); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(influencer.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to hash password",
		})
	}

	influencer.Password = string(hashedPassword)

	// get jwt from id returned by AddInfluencer
	if _, err := cn.Controller.AddInfluencer(*influencer); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, "success register")
}
