package controller

import (
	"encoding/json"
	"influence-hub-influencer/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cn *Controller) ShowCampaign(c echo.Context) error {
	url := "placeholderUrlBrand"

	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error":   err.Error(),
			"details": "error at sending request to brand server",
		})
		return err
	}
	defer res.Body.Close()
	var campaigns []models.Campaign
	err = json.NewDecoder(res.Body).Decode(&campaigns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error":   err.Error(),
			"details": "error at decoding response from brand server",
		})
		return err
	}

	return c.JSON(http.StatusOK, campaigns)
}

func (cn *Controller) ApplyCampaign(c echo.Context) error {

	return nil
}
