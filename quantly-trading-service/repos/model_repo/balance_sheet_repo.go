package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type BalanceSheetRepo struct {
	db *gorm.DB
}

func (r *BalanceSheetRepo) GetById(id string) (*models.BalanceSheet, error) {
	var balanceSheet models.BalanceSheet
	if err := r.db.First(&balanceSheet, "stock_symbol = ?", id).Error; err != nil {
		return nil, utils.OnError(err, "can not get balance sheet")
	}

	return &balanceSheet, nil
}

func (r *BalanceSheetRepo) Create(balanceSheet *models.BalanceSheet) error {
	err := r.db.Create(balanceSheet).Error
	return utils.OnError(err, "can not create balance sheet")
}

func (r *BalanceSheetRepo) Update(balanceSheet *models.BalanceSheet) error {
	if balanceSheet.StockSymbol == "" {
		return fmt.Errorf("input invalid")
	}
	err := r.db.Save(balanceSheet).Error
	return utils.OnError(err, "can not update balane sheet")
}

func (r *BalanceSheetRepo) DeleteById(id string) error {
	err := r.db.Where("stock_symbol = ?", id).Delete(&models.BalanceSheet{}).Error
	return utils.OnError(err, "can not delete balance sheet")
}
