package middleware

import (
	"influence-hub-influencer/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct {
	Authentication repository.Repository
}

func NewAuth(repo repository.Repository) *Auth {
	return &Auth{Authentication: repo}
}

func (a *Auth) AuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")

		claims, err := ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "wrong token",
			})
		}

		// Mencoba mengonversi ID dari string ke primitive.ObjectID
		id, err := primitive.ObjectIDFromHex(claims["id"].(string))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid value token",
			})
		}

		loggedinInfluencer, err := a.Authentication.FindById(id)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid value token",
			})
		}
		c.Set("loggedinInfluencer", loggedinInfluencer.ID)
		return next(c)
	}
}
