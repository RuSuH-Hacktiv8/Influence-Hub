package controller

import (
	"bytes"
	"encoding/json"
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
			"error":   err.Error(),
		})
	}

	influencer.Password = string(hashedPassword)

	followers, err := middleware.GetInstagramFollowers(influencer.InstagramUsername)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to fetch Instagram followers",
			"error":   err.Error(),
		})
	}

	influencer.InstagramFollowers = followers

	resultID, err := cn.Controller.AddInfluencer(*influencer)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to register",
			"error":   err.Error(),
		})
	}

	secretKey := os.Getenv("SECRET_KEY")

	influencer.ID = resultID

	token, err := middleware.GenerateJWT(influencer, secretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to generate JWT",
			"error":   err.Error(),
		})
	}
	if err := registerNotificationRequest(c, influencer); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed sending email notification",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Register successful",
		"token":   token,
	})
}

func registerNotificationRequest(c echo.Context, influencer *models.Influencer) error {
	// change the url when deploying on gcp
	url := "http://localhost:8082"
	endpoint := "/mails/success_register"
	j, err := json.Marshal(map[string]any{
		"username": influencer.Name,
		"email":    influencer.Email,
		"role":     "influencer",
	})
	if err != nil {
		return err
	}
	payload := bytes.NewBuffer(j)
	req, _ := http.NewRequest("POST", url+endpoint, payload)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error":   err.Error(),
			"details": "error at sending request to brand server",
		})
		return err
	}
	defer res.Body.Close()
	return nil
}
