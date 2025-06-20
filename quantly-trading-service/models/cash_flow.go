package models

type CashFlow struct {
	StockSymbol                   string  `json:"stock_symbol" gorm:"primaryKey;type:varchar(64)"`
	OperatingCashFlow             float64 `json:"operating_cash_flow"`
	InvestingCashFlow             float64 `json:"investing_cash_flow"`
	FinancingCashFlow             float64 `json:"financing_cash_flow"`
	EndCashPosition               float64 `json:"end_cash_position"`
	IncomeTaxPaidSupplementalData float64 `json:"income_tax_paid_supplemental_data"`
	InterestPaidSupplementalData  float64 `json:"interest_paid_supplemental_data"`
	CapitalExpenditure            float64 `json:"capital_expenditure"`
	IssuanceOfDebt                float64 `json:"issuance_of_debt"`
	RepaymentOfDebt               float64 `json:"repayment_of_debt"`
	RepurchaseOfCapitalStock      float64 `json:"repurchase_of_capital_stock"`
	FreeCashFlow                  float64 `json:"free_cash_flow"`
}
