package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type StockRepo struct {
	db *gorm.DB
}

func (r *StockRepo) GetById(stockSymbol string) (*models.Stock, error) {
	var stock *models.Stock
	err := r.db.First(stock, "stock_symbol = ?", stockSymbol).Error
	if err != nil {
		return nil, fmt.Errorf("can not get stock by id: %w", err)
	}
	return stock, nil
}

func (r *StockRepo) Create(stock *models.Stock) error {
	err := r.db.Create(stock).Error
	if err != nil {
		return fmt.Errorf("can not create new stock: %w", err)
	}
	return nil
}

func (r *StockRepo) Update(stock *models.Stock) error {
	if stock.Symbol == "" {
		return fmt.Errorf("can not update stock: invalid input")
	}

	err := r.db.Save(stock).Error
	if err != nil {
		return fmt.Errorf("can not update stock: %w", err)
	}
	return nil
}

func (r *StockRepo) DeleteById(stockSymbol string) error {
	err := r.db.Where("stock_symbol = ?", stockSymbol).Delete(&models.Stock{}).Error
	if err != nil {
		return fmt.Errorf("can not delete stock: %w", err)
	}
	return nil
}
