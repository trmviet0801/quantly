package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type CashFlowRepo struct {
	DB *gorm.DB
}

func (r *CashFlowRepo) GetById(stockSymbol string) (*models.CashFlow, error) {
	cashFlow := &models.CashFlow{}
	if err := r.DB.First(cashFlow, "stock_symbol = ?", stockSymbol).Error; err != nil {
		return nil, utils.OnError(err, "can not get cash flow")
	}
	return cashFlow, nil
}

func (r *CashFlowRepo) Create(cashFlow *models.CashFlow) error {
	err := r.DB.Create(cashFlow).Error
	return utils.OnError(err, "can not create cash flow")
}

func (r *CashFlowRepo) Update(cashFlow *models.CashFlow) error {
	if cashFlow.StockSymbol == "" {
		return fmt.Errorf("can not update cashflow: input invalid")
	}

	err := r.DB.Save(cashFlow).Error
	return utils.OnError(err, "can not update cash flow")
}

func (r *CashFlowRepo) DeleteById(stockSymbol string) error {
	err := r.DB.Where("stock_symbol = ?", stockSymbol).Delete(&models.CashFlow{}).Error
	return utils.OnError(err, "can not delete cash flow")
}
