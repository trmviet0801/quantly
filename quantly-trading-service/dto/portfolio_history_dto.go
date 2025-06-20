package dto

type PortfolioHistoryDto struct {
	Timestamp     []int64   `json:"timestamp"`
	Equity        []float64 `json:"equity"`
	ProfitLoss    []float64 `json:"profit_loss"`
	ProfitLossPct []float64 `json:"profit_loss_pct"`
	BaseValue     float64   `json:"base_value"`
	BaseValueAsof string    `json:"base_value_asof"`
	Timeframe     string    `json:"timeframe"`
}
