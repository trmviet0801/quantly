package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type CashFlowRepo struct {
	db *gorm.DB
}

func (r *CashFlowRepo) GetById(stockSymbol string) (*models.CashFlow, error) {
	var cashFlow *models.CashFlow
	err := r.db.First(cashFlow, "stock_symbol = ?", stockSymbol).Error
	if err != nil {
		return nil, fmt.Errorf("can not get cash flow: %w", err)
	}
	return cashFlow, nil
}

func (r *CashFlowRepo) Create(cashFlow *models.CashFlow) error {
	err := r.db.Create(cashFlow).Error
	if err != nil {
		return fmt.Errorf("can not create cash flow: %w", err)
	}
	return nil
}

func (r *CashFlowRepo) Update(cashFlow *models.CashFlow) error {
	if cashFlow.StockSymbol == "" {
		return fmt.Errorf("can not update cashflow: input invalid")
	}
	err := r.db.Save(cashFlow).Error
	if err != nil {
		return fmt.Errorf("can not update cash flow: %w", err)
	}
	return nil
}

func (r *CashFlowRepo) DeleteById(stockSymbol string) error {
	err := r.db.Where("stock_symbol = ?", stockSymbol).Delete(&models.CashFlow{}).Error
	if err != nil {
		return fmt.Errorf("can not delete cash flow: %w", err)
	}
	return nil
}
