package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type StockRepo struct {
	db *gorm.DB
}

func (r *StockRepo) GetById(stockSymbol string) (*models.Stock, error) {
	stock := &models.Stock{}
	if err := r.db.First(stock, "stock_symbol = ?", stockSymbol).Error; err != nil {
		return nil, utils.OnError(err, "can not get stock")
	}
	return stock, nil
}

func (r *StockRepo) Create(stock *models.Stock) error {
	err := r.db.Create(stock).Error
	return utils.OnError(err, "can not create new stock")
}

func (r *StockRepo) Update(stock *models.Stock) error {
	if stock.Symbol == "" {
		return fmt.Errorf("can not update stock: invalid input")
	}

	err := r.db.Save(stock).Error
	return utils.OnError(err, "can not update stock")
}

func (r *StockRepo) DeleteById(stockSymbol string) error {
	err := r.db.Where("stock_symbol = ?", stockSymbol).Delete(&models.Stock{}).Error
	return utils.OnError(err, "can not delete stock")
}
