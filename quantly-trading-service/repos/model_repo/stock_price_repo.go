package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type StockPriceRepo struct {
	DB *gorm.DB
}

func (r *StockPriceRepo) GetById(symbol string) (*models.StockPrice, error) {
	stockPrice := &models.StockPrice{}
	if err := r.DB.First(&stockPrice, "symbol = ?", symbol).Error; err != nil {
		return nil, err
	}
	return stockPrice, nil
}

func (r *StockPriceRepo) Create(stockPrice *models.StockPrice) error {
	err := r.DB.Create(stockPrice).Error
	return utils.OnError(err, "can not create stock price")
}

func (r *StockPriceRepo) Update(stockPrice *models.StockPrice) error {
	if stockPrice.Symbol == "" {
		return fmt.Errorf("input not valid")
	}

	err := r.DB.Save(stockPrice).Error
	return utils.OnError(err, "can not save stock price")
}

func (r *StockPriceRepo) DeleteById(symbol string) error {
	err := r.DB.Where("symbol = ?", symbol).Delete(&models.StockPrice{}).Error
	return utils.OnError(err, "can not delete stock price")
}
