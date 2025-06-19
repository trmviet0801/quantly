package models

import "fmt"

type Position struct {
	PositionId             string  `json:"position_id" gorm:"primaryKey"`
	AccountId              string  `json:"account_id"`
	AssetId                string  `json:"asset_id"`
	Symbol                 string  `json:"symbol"`
	Exchange               string  `json:"exchange"`
	AssetClass             string  `json:"asset_class"`
	AssetMarginable        bool    `json:"asset_marginable"`
	Quantity               float64 `json:"qty"`
	AverageEntryPrice      float64 `json:"avg_entry_price"`
	Side                   string  `json:"side"`
	MarketValue            float64 `json:"market_value"`
	CostBasis              float64 `json:"cost_basis"`
	UnrealizedPL           float64 `json:"unrealized_pl"`
	UnrealizedPLPC         float64 `json:"unrealized_plpc"`
	UnrealizedIntradayPL   float64 `json:"unrealized_intraday_pl"`
	UnrealizedIntradayPLPC float64 `json:"unrealized_intraday_plpc"`
	CurrentPrice           float64 `json:"current_price"`
	LastDayPrice           float64 `json:"lastday_price"`
	ChangeToday            float64 `json:"change_today"`
	QuantityAvailable      float64 `json:"qty_available"`
}

func (p *Position) String() string {
	return fmt.Sprintf(
		"Position{ID:%s, AccountID:%s, Symbol:%s, Side:%s, Qty:%.2f, AvgPrice:%.2f, CurrentPrice:%.2f, MarketValue:%.2f, UnrealizedPL:%.2f (%.2f%%), IntradayPL:%.2f (%.2f%%)}",
		p.PositionId,
		p.AccountId,
		p.Symbol,
		p.Side,
		p.Quantity,
		p.AverageEntryPrice,
		p.CurrentPrice,
		p.MarketValue,
		p.UnrealizedPL,
		p.UnrealizedPLPC*100,
		p.UnrealizedIntradayPL,
		p.UnrealizedIntradayPLPC*100,
	)
}
