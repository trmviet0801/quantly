package models

type Position struct {
	PositionId             int64   `json:"position_id"`
	AccountId              int64   `json:"account_id"`
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
