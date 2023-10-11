package controller

import (
	"influence-hub-influencer/middleware"
	"influence-hub-influencer/models"
	"net/http"
	"os"

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

	followers, err := middleware.GetInstagramFollowers(influencer.InstagramUsername)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to fetch Instagram followers",
		})
	}

	influencer.InstagramFollowers = followers

	resultID, err := cn.Controller.AddInfluencer(*influencer)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to register",
		})
	}

	secretKey := os.Getenv("SECRET_KEY")

	influencer.ID = resultID

	token, err := middleware.GenerateJWT(influencer, secretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to generate JWT",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Register successful",
		"token":   token,
	})
}
