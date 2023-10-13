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

func (cc CampaignController) CreateCampaign(c echo.Context) error {
	// Mengambil brandID dari klaim token
	brandID := c.Get("loggedInBrand").(uint)

	campaign := new(models.Campaign)

	if err := c.Bind(campaign); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Set brandID pada campaign
	campaign.BrandID = brandID

	newCampaign, err := cc.Repo.AddCampaign(*campaign)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, newCampaign)
}

func (cc CampaignController) UpdateCampaign(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	campaign := new(models.Campaign)

	if err := c.Bind(campaign); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	result, err := cc.Repo.EditCampaign(uint(id), *campaign)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)

}

func (cc CampaignController) DeleteCampaign(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := cc.Repo.DeletesCampaign(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, echo.Map{
		"message": "success delete campaign",
	})
}

func (cc *CampaignController) GetCampaign(e echo.Context) error {
	id := e.Get("loggedInBrand").(string)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid campaign ID",
		})
	}

	campaigns, err := cc.Repo.GetCampaign(idInt)
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

func (cc *CampaignController) GetAllCampaign(e echo.Context) error {
	// Panggil metode GetAllCampaign dari repository
	campaigns, err := cc.Repo.GetAllCampaign()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusOK, campaigns)
}
