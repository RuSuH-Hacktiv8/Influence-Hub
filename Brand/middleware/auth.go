package middleware

import (
	"influence-hub-brand/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Auth struct {
	Authentication repository.Repository
}

func NewAuth(repo repository.Repository) *Auth {
	return &Auth{Authentication: repo}
}

func (a *Auth) AuthBrand(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")

		claims, err := ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "wrong token",
			})
		}

		id := claims["id"].(float64)
		loggedInBrand, err := a.Authentication.FindById(int(id))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid value token",
			})
		}
		c.Set("loggedInBrand", loggedInBrand.ID)
		return next(c)
	}

}
