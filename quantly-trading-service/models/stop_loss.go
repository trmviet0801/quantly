package models

type StopLoss struct {
	StopLostId int64   `json:"stop_loss_id" binding:"required" gorm:"primaryKey"`
	OrderId    int64   `json:"order_id" binding:"required"`
	StopPrice  float64 `json:"stop_price" binding:"required"`
	LimitPrice float64 `json:"limit_price" binding:"required"`
}
