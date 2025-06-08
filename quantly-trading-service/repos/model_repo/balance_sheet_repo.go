package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type BalanceSheetRepo struct {
	db *gorm.DB
}

func (r *BalanceSheetRepo) GetById(id string) (*models.BalanceSheet, error) {
	var balanceSheet models.BalanceSheet
	err := r.db.First(&balanceSheet, "stock_symbol = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("can not find balance sheet %w", err)
	}
	return &balanceSheet, nil
}

func (r *BalanceSheetRepo) Create(balanceSheet *models.BalanceSheet) error {
	err := r.db.Create(balanceSheet).Error
	if err != nil {
		return fmt.Errorf("can not create balance sheet %w", err)
	}
	return nil
}

func (r *BalanceSheetRepo) Update(balanceSheet *models.BalanceSheet) error {
	if balanceSheet.StockSymbol == "" {
		return fmt.Errorf("input invalid")
	}
	err := r.db.Save(balanceSheet).Error
	if err != nil {
		return fmt.Errorf("can not update balane sheet: %w", err)
	}
	return nil
}

func (r *BalanceSheetRepo) DeleteById(id string) error {
	err := r.db.Where("stock_symbol = ?", id).Delete(&models.BalanceSheet{}).Error
	if err != nil {
		return fmt.Errorf("can not delete balance sheet: %w", err)
	}
	return nil
}
