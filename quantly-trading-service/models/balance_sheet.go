package models

type BalanceSheet struct {
	StockSymbol                         string  `json:"stock_symbol" gorm:"primaryKey;type:varchar(64)"`
	TotalAssets                         float64 `json:"total_assets"`
	TotalLiabilitiesNetMinorityInterest float64 `json:"total_liabilities_net_minority_interest"`
	TotalEquityGrossMinorityInterest    float64 `json:"total_equity_gross_minority_interest"`
	TotalCapitalization                 float64 `json:"total_capitalization"`
	CommonStockEquity                   float64 `json:"common_stock_equity"`
	CapitalLeaseObligations             float64 `json:"capital_lease_obligations"`
	NetTangibleAssets                   float64 `json:"net_tangible_assets"`
	WorkingCapital                      float64 `json:"working_capital"`
	InvestedCapital                     float64 `json:"invested_capital"`
	TangibleBookValue                   float64 `json:"tangible_book_value"`
	TotalDebt                           float64 `json:"total_debt"`
	NetDebt                             float64 `json:"net_debt"`
	ShareIssued                         float64 `json:"share_issued"`
	OrdinarySharesNumber                float64 `json:"ordinary_shares_number"`
}
