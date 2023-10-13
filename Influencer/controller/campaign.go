package controller

import (
	"bytes"
	"encoding/json"
	"influence-hub-influencer/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cn *Controller) ShowCampaign(c echo.Context) error {
	// Ganti "placeholderUrlBrand" dengan URL yang sesuai
	url := "http://localhost:8081/campaign"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err.Error(),
			"details": "error at creating request to brand server",
		})
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err.Error(),
			"details": "error at sending request to brand server",
		})
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.JSON(res.StatusCode, map[string]interface{}{
			"error":   "Brand server returned a non-200 status code",
			"details": res.Status,
		})
		return nil
	}

	var campaigns []models.Campaign
	err = json.NewDecoder(res.Body).Decode(&campaigns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err.Error(),
			"details": "error at decoding response from brand server",
		})
		return err
	}

	return c.JSON(http.StatusOK, campaigns)
}

func (cn *Controller) ApplyCampaign(c echo.Context) error {
	// Deklarasikan variabel untuk objek Contract
	var contract models.Contract

	// Bind data JSON dari body permintaan ke objek contract
	if err := c.Bind(&contract); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to bind request data"})
	}

	// Marshal objek contract ke JSON
	contractData, err := json.Marshal(contract)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal contract data"})
	}

	// Membuat permintaan HTTP POST ke endpoint localhost:8081/contract
	resp, err := http.Post("http://localhost:8081/contract", "application/json", bytes.NewBuffer(contractData))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	if err := ApplyNotificationRequest(c, &contract); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed sending email notification",
			"error":   err.Error(),
		})
	}

	// Jika berhasil, Anda dapat mengembalikan respons sukses
	return c.JSON(http.StatusOK, map[string]string{"message": "Campaign applied successfully"})
}

func ApplyNotificationRequest(c echo.Context, contract *models.Contract) error {
	// change the url when deploying on gcp
	url := "http://localhost:8082"
	endpoint := "/mails/success_apply_campaign"
	j, err := json.Marshal(map[string]any{
		"ID":           contract.ID,
		"CampaignID":   contract.CampaignID,
		"InfluencerID": contract.InfluencerID,
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
