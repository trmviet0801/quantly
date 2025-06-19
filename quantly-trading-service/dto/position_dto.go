package dto

type PositionDto struct {
	AssetId                string `json:"asset_id"`
	Symbol                 string `json:"symbol"`
	Exchange               string `json:"exchange"`
	AssetClass             string `json:"asset_class"`
	AssetMarginable        bool   `json:"asset_marginable"`
	Quantity               string `json:"qty"`
	AverageEntryPrice      string `json:"avg_entry_price"`
	Side                   string `json:"side"`
	MarketValue            string `json:"market_value"`
	CostBasis              string `json:"cost_basis"`
	UnrealizedPL           string `json:"unrealized_pl"`
	UnrealizedPLPC         string `json:"unrealized_plpc"`
	UnrealizedIntradayPL   string `json:"unrealized_intraday_pl"`
	UnrealizedIntradayPLPC string `json:"unrealized_intraday_plpc"`
	CurrentPrice           string `json:"current_price"`
	LastDayPrice           string `json:"lastday_price"`
	ChangeToday            string `json:"change_today"`
	QuantityAvailable      string `json:"qty_available"`
}
