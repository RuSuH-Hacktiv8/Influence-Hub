package controller

import (
	"influence-hub-brand/models"
	"influence-hub-brand/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CampaignController struct {
	Repo repository.Repository
}

func NewCampaignController(repo repository.Repository) CampaignController {
	return CampaignController{repo}
}

func (c *CampaignController) AddCampaign(e echo.Context) error {
	campaign := new(models.Campaign)
	if err := e.Bind(campaign); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid request")
	}

	campaigns, err := c.Repo.AddCampaign(*campaign)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to add campaign",
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"message":   "Campaign added successfully",
		"campaigns": campaigns,
	})

}

func (c *CampaignController) GetCampaign(e echo.Context) error {
	id := e.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid campaign ID",
		})
	}

	campaigns, err := c.Repo.GetCampaign(idInt)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get campaigns",
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"message":   "Campaigns retrieved successfully",
		"campaigns": campaigns,
	})
}
