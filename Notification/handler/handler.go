package handler

import (
	"fmt"
	"influence-hub-notification/models"
	"influence-hub-notification/notification"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SuccessRegister(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error":   err.Error(),
			"details": "Failed on binding json input",
		})
		return err
	}

	if err := notification.SuccessRegisterEmail(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error":   err.Error(),
			"details": "Failed on sending email",
		})
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("email Sent to %s", user.Email),
	})
}

func SuccessApplyCampaign(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error":   err.Error(),
			"details": "Failed on binding json input",
		})
		return err
	}

	if err := notification.SuccessApplyCampaign(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error":   err.Error(),
			"details": "Failed on sending email",
		})
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("email Sent to %s", user.Email),
	})
}

func PaymentReceived(c echo.Context) error {

	return nil
}
