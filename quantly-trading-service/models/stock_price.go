package models

import "time"

type StockPrice struct {
	Symbol    string    `gorm:"column:symbol;primaryKey" binding:"required"`
	Price     float64   `gorm:"column:price" binding:"required"`
	Timestamp time.Time `gorm:"column:timestamp;primaryKey" binding:"required"`
}
