package dto

import "fmt"

type OrderResponseDto struct {
	AssetID                string `json:"asset_id"`
	Symbol                 string `json:"symbol"`
	Exchange               string `json:"exchange"`
	AssetClass             string `json:"asset_class"`
	AssetMarginable        bool   `json:"asset_marginable"`
	Qty                    string `json:"qty"`
	AvgEntryPrice          string `json:"avg_entry_price"`
	Side                   string `json:"side"`
	MarketValue            string `json:"market_value"`
	CostBasis              string `json:"cost_basis"`
	UnrealizedPL           string `json:"unrealized_pl"`
	UnrealizedPLPC         string `json:"unrealized_plpc"`
	UnrealizedIntradayPL   string `json:"unrealized_intraday_pl"`
	UnrealizedIntradayPLPC string `json:"unrealized_intraday_plpc"`
	CurrentPrice           string `json:"current_price"`
	LastdayPrice           string `json:"lastday_price"`
	ChangeToday            string `json:"change_today"`
	QtyAvailable           string `json:"qty_available"`
}

func (o *OrderResponseDto) String() string {
	return fmt.Sprintf(
		"OrderResponseDto{\n"+
			"  AssetID: %s\n"+
			"  Symbol: %s\n"+
			"  Exchange: %s\n"+
			"  AssetClass: %s\n"+
			"  AssetMarginable: %t\n"+
			"  Qty: %s\n"+
			"  AvgEntryPrice: %s\n"+
			"  Side: %s\n"+
			"  MarketValue: %s\n"+
			"  CostBasis: %s\n"+
			"  UnrealizedPL: %s\n"+
			"  UnrealizedPLPC: %s\n"+
			"  UnrealizedIntradayPL: %s\n"+
			"  UnrealizedIntradayPLPC: %s\n"+
			"  CurrentPrice: %s\n"+
			"  LastdayPrice: %s\n"+
			"  ChangeToday: %s\n"+
			"  QtyAvailable: %s\n"+
			"}",
		o.AssetID,
		o.Symbol,
		o.Exchange,
		o.AssetClass,
		o.AssetMarginable,
		o.Qty,
		o.AvgEntryPrice,
		o.Side,
		o.MarketValue,
		o.CostBasis,
		o.UnrealizedPL,
		o.UnrealizedPLPC,
		o.UnrealizedIntradayPL,
		o.UnrealizedIntradayPLPC,
		o.CurrentPrice,
		o.LastdayPrice,
		o.ChangeToday,
		o.QtyAvailable,
	)
}
