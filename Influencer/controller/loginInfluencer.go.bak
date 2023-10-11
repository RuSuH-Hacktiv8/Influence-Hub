package controller

import (
	"influence-hub-influencer/middleware"
	"influence-hub-influencer/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (cn *Controller) Login(c echo.Context) error {
	influencer := new(models.Influencer)
	if err := c.Bind(influencer); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Fetch the influencer from the database based on the email (or username) provided in the request.
	fetchedInfluencer, err := cn.Controller.FindByEmail(influencer.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "User not found",
		})
	}

	// Compare the hashed password from the database with the provided password.
	err = bcrypt.CompareHashAndPassword([]byte(fetchedInfluencer.Password), []byte(influencer.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Incorrect password",
		})
	}

	// If the passwords match, generate a JWT for the user and return it in the response.
	token, err := middleware.GenerateJWT(fetchedInfluencer, "secret")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to generate JWT",
		})
	}

	// You can return the JWT in the response for the client to use in subsequent requests.
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"token":   token,
	})
}
