package models

type Order struct {
	OrderId       int64   `json:"order_id" binding:"required" gorm:"primaryKey"`
	AccountId     int64   `json:"account_id" binding:"required"`
	Symbol        string  `json:"symbol" binding:"required"`
	Quantity      float64 `json:"qty" binding:"required"`
	Notional      float64 `json:"notional"`
	Side          string  `json:"side"`
	Type          string  `json:"type"`
	TimeInForce   string  `json:"time_in_force"`
	LimitPrice    float64 `json:"limit_price"`
	StopPrice     float64 `json:"stop_price"`
	TrailPrice    float64 `json:"trail_price"`
	TrailPercent  float64 `json:"trail_percent"`
	ExtendedHours bool    `json:"extended_hours"`
	ClientOrderId string  `json:"client_order_id"`
	OrderClass    string  `json:"order_class"`
}
