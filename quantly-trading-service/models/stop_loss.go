package models

type StopLoss struct {
	StopLosstId string  `json:"stop_loss_id" binding:"required" gorm:"primaryKey;type:varchar(64)"`
	OrderId     string  `json:"order_id" binding:"required" gorm:"type:varchar(64)"`
	StopPrice   float64 `json:"stop_price" binding:"required"`
	LimitPrice  float64 `json:"limit_price" binding:"required"`
}
