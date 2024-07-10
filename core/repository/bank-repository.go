package repository

import (
	"github.com/akposieyefa/open-ai/core/models"
	"github.com/akposieyefa/open-ai/core/types"
	"github.com/akposieyefa/open-ai/internals"
)

func GetBanks() ([]types.BankResponse, error) {

	var banks []models.Bank

	if err := internals.DB.First(&banks).Error; err != nil {
		return nil, err
	}

	bankResponse := make([]types.BankResponse, 0)

	for _, bank := range banks {
		bankResponse = append(bankResponse, types.BankResponse{
			ID:        bank.ID,
			Name:      bank.Name,
			Code:      bank.Code,
			Ussd:      bank.Ussd,
			Logo:      bank.Logo,
			Slug:      bank.Slug,
			CreatedAt: bank.CreatedAt,
			UpdatedAt: bank.UpdatedAt,
			DeletedAt: bank.DeletedAt,
		})
	}

	return bankResponse, nil

}
