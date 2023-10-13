package controller

import (
	"influence-hub-brand/models"
	"influence-hub-brand/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContractController struct {
	Repo repository.Repository
}

func NewContractController(repo repository.Repository) ContractController {
	return ContractController{repo}
}

func (cc ContractController) AddContract(c echo.Context) error {
	// Mendeklarasikan variabel untuk menyimpan data kontrak dari permintaan JSON
	var contract models.Contract

	// Mengikat data JSON dari permintaan ke variabel contract
	if err := c.Bind(&contract); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to bind request data"})
	}

	// Memanggil repository untuk menambahkan kontrak
	contractID, err := cc.Repo.AddContract(contract)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Mengembalikan ID kontrak yang telah ditambahkan
	return c.JSON(http.StatusCreated, map[string]uint{"contract_id": contractID})
}
