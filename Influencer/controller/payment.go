package controller

import (
	"influence-hub-influencer/middleware"
	"influence-hub-influencer/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cn *Controller) RequestPayment(c echo.Context) error {
    // Deklarasikan variabel untuk objek PaymentData
    var paymentData models.PaymentData

    // Bind data JSON dari body permintaan ke objek paymentData
    if err := c.Bind(&paymentData); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to bind request data"})
    }

    // Selanjutnya, Anda dapat mengirim data paymentData ke middleware untuk melakukan permintaan pembayaran.

    err := middleware.SendPaymentRequest(paymentData)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Payment request sent"})
}

