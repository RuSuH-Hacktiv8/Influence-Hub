package controller

import (
	"bytes"
	"encoding/json"
	"influence-hub-brand/middleware"
	"influence-hub-brand/models"
	"influence-hub-brand/repository"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type BrandController struct {
	Repo repository.Repository
}

func NewBrandController(repo repository.Repository) BrandController {
	return BrandController{repo}
}

func (bc BrandController) Register(c echo.Context) error {
	brand := new(models.Brand)
	if err := c.Bind(brand); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(brand.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to hash password",
		})
	}

	brand.Password = string(hashedPassword)

	// get jwt from id returned by AddBrand
	resultID, err := bc.Repo.AddBrand(*brand)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to register",
		})
	}

	secretKey := os.Getenv("SECRET_KEY")

	brand.ID = resultID
	token, err := middleware.GenerateJWT(brand, secretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to generate JWT",
		})
	}

	if err := registerNotificationRequest(c, brand); err != nil {
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

func (bc BrandController) Login(c echo.Context) error {
	brand := new(models.Brand)
	if err := c.Bind(brand); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Fetch the brand from the database based on the email (or username) provided in the request.
	fetchedBrand, err := bc.Repo.FindByEmail(brand.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "User not found",
		})
	}

	// Compare the hashed password from the database with the provided password.
	err = bcrypt.CompareHashAndPassword([]byte(fetchedBrand.Password), []byte(brand.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Incorrect password",
		})
	}

	// If the passwords match, generate a JWT for the user and return it in the response.
	token, err := middleware.GenerateJWT(&fetchedBrand, "secret")
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

func registerNotificationRequest(c echo.Context, brand *models.Brand) error {
	// change the url when deploying on gcp
	url := "http://localhost:8082"
	endpoint := "/mails/success_register"
	j, err := json.Marshal(map[string]any{
		"username": brand.Name,
		"email":    brand.Email,
		"role":     "brand",
	})
	if err != nil {
		return err
	}
	payload := bytes.NewBuffer(j)
	req, _ := http.NewRequest("POST", url+endpoint, payload)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error":   err.Error(),
			"details": "error at sending request to notif server",
		})
		return err
	}
	defer res.Body.Close()
	var responseMap map[string]any
	if err := json.NewDecoder(res.Body).Decode(&responseMap); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, responseMap)
}
