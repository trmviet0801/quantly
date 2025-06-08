package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type IncomeRepo struct {
	db *gorm.DB
}

func (r *IncomeRepo) GetById(stockSymbol string) (*models.Income, error) {
	var income *models.Income
	err := r.db.First(income, "stock_symbol = ?", stockSymbol).Error
	if err != nil {
		return nil, fmt.Errorf("can not find income: %w", err)
	}
	return income, nil
}

func (r *IncomeRepo) Create(income *models.Income) error {
	err := r.db.Create(income).Error
	if err != nil {
		return fmt.Errorf("can not create income: %w", err)
	}
	return nil
}

func (r *IncomeRepo) Update(income *models.Income) error {
	if income.StockSymbol == "" {
		return fmt.Errorf("can not update income: input invalid")
	}

	err := r.db.Save(income).Error
	if err != nil {
		return fmt.Errorf("can not update income: %w", err)
	}
	return nil
}

func (r *IncomeRepo) DeleteById(stockSymbol string) error {
	err := r.db.Where("stock_symbol = ?", stockSymbol).Delete(&models.Income{}).Error
	if err != nil {
		return fmt.Errorf("can not delete: %w", err)
	}
	return nil
}
