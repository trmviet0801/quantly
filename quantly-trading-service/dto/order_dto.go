package dto

type OrderDto struct {
	Symbol      string `json:"symbol"`
	Qty         string `json:"qty"`
	Side        string `json:"side"`
	Type        string `json:"type"`
	TimeInForce string `json:"time_in_force"`
	LimitPrice  string `json:"limit_price"`
	// StopPrice      string `json:"stop_price"`]
	OrderClass    string `json:"order_class"`
	ExtendedHours bool   `json:"extended_hours"`
	// ClientOrderId  string `json:"client_order_id"`
	// Commission     string `json:"commission"`
	// CommissionType string `json:"commission_type"`
}

func (orderDto *OrderDto) IsValid() bool {
	return (orderDto.Type != "" &&
		orderDto.TimeInForce != "" &&
		orderDto.Side != "" &&
		orderDto.LimitPrice != "")
}
