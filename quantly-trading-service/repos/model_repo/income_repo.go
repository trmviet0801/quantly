package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type IncomeRepo struct {
	db *gorm.DB
}

func (r *IncomeRepo) GetById(stockSymbol string) (*models.Income, error) {
	income := &models.Income{}
	if err := r.db.First(income, "stock_symbol = ?", stockSymbol).Error; err != nil {
		return nil, utils.OnError(err, "can not find income")
	}
	return income, nil
}

func (r *IncomeRepo) Create(income *models.Income) error {
	err := r.db.Create(income).Error
	return utils.OnError(err, "can not create income")
}

func (r *IncomeRepo) Update(income *models.Income) error {
	if income.StockSymbol == "" {
		return fmt.Errorf("can not update income: input invalid")
	}

	err := r.db.Save(income).Error
	return utils.OnError(err, "can not update income")
}

func (r *IncomeRepo) DeleteById(stockSymbol string) error {
	err := r.db.Where("stock_symbol = ?", stockSymbol).Delete(&models.Income{}).Error
	return utils.OnError(err, "can not delete income")
}
