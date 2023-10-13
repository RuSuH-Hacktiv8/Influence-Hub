package repository

import "influence-hub-brand/models"

func (r Repository) AddContract(contract models.Contract) (uint, error) {
    query := r.DB.Create(&contract)
    if query.Error != nil {
        return 0, query.Error // Mengembalikan 0 dan pesan kesalahan
    }

    return contract.ID, nil
}
