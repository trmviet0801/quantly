package models

type TakeProfit struct {
	TakeProfitId int64   `gorm:"primaryKey" json:"take_profit_id" binding:"required"`
	OrderId      int64   `json:"order_id" binding:"required" gorm:"foreignKey:OrderId;references:OrderId"`
	LimitPrice   float64 `json:"limit_price" binding:"required"`
}
