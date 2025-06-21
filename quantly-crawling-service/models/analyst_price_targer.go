package models

type AnalystPriceTarget struct {
	Low     float64 `json:"low" redis:"low"`
	Current float64 `json:"current" redis:"current"`
	Average float64 `json:"average" redis:"average"`
	High    float64 `json:"high" redis:"high"`
}
