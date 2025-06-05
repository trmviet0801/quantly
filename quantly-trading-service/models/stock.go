package models

type Stock struct {
	Name            string  `binding:"required" json:"name"`
	Symbol          string  `binding:"required" json:"symbol"`
	IpoYear         int16   `json:"ipo_year"`
	Country         string  `binding:"required" json:"country"`
	CurrentPrice    float64 `json:"current_price"`
	PriceChange     float64 `json:"price_change"`
	ChangePercent   float32 `json:"change_percent"`
	OpenPrice       float64 `json:"open_price"`
	DayRange        float64 `json:"day_range"`
	DayLow          float64 `json:"day_low"`
	DayHigh         float64 `json:"day_high"`
	Volume          int64   `json:"volume"`
	LatestTradeTime string  `json:"latest_trade_time"`
	Ticker          string  `json:"ticker"`
	Exchange        string  `json:"exchange"`
	Industry        string  `binding:"required" json:"industry"`
	Sector          string  `binding:"required" json:"sector"`
	Employees       int32   `json:"employees"`
	Headquarters    string  `json:"headquarters"`
	MarketCap       float64 `json:"market_cap"`
	PERatioTtm      float32 `json:"pe_ratio_ttm"`
	EPSTtm          float32 `json:"eps_ttm"`
}
