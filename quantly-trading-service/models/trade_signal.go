package models

import "time"

type TradeSignal struct {
	TradeSignalId int64     `json:"trade_signal_id" gorm:"primaryKey;autoIncrement"`
	Type          string    `json:"type" gorm:"not null"`
	StockSymbol   string    `json:"stock_symbol" gorm:"not null"`
	Quant         string    `json:"quant" gorm:"not null"`
	TimeStamp     time.Time `json:"timestamp" gorm:"not null"`
	QuantModelId  int64     `json:"quant_model_id" gorm:"foreignKey:QuantModelId;references:QuantModelId;not null"`
	AccountId     int64     `json:"account_id" gorm:"foreignKey:AccountId;references:AccountId;not null"`
}
