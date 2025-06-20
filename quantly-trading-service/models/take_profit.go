package models

type TakeProfit struct {
	TakeProfitId string  `gorm:"primaryKey;type:varchar(64)" json:"take_profit_id" binding:"required"`
	OrderId      string  `json:"order_id" binding:"required" gorm:"foreignKey:OrderId;references:OrderId;type:varchar(64)"`
	LimitPrice   float64 `json:"limit_price" binding:"required"`
}
