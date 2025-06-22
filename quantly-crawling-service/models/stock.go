package models

import (
	"encoding/json"
	"fmt"
)

type Stock struct {
	Name                              string                 `json:"name" redis:"name"`
	CompanyID                         string                 `json:"company_id" redis:"company_id"`
	EntityType                        string                 `json:"entity_type" redis:"entity_type"`
	Summary                           string                 `json:"summary" redis:"summary"`
	StockTicker                       string                 `json:"stock_ticker" redis:"stock_ticker"`
	Currency                          string                 `json:"currency" redis:"currency"`
	EarningsDate                      string                 `json:"earnings_date" redis:"earnings_date"`
	Exchange                          string                 `json:"exchange" redis:"exchange"`
	ClosingPrice                      float64                `json:"closing_price" redis:"closing_price"`
	PreviousClose                     float64                `json:"previous_close" redis:"previous_close"`
	Open                              float64                `json:"open" redis:"open"`
	Bid                               string                 `json:"bid" redis:"bid"`
	Ask                               string                 `json:"ask" redis:"ask"`
	DayRange                          string                 `json:"day_range" redis:"day_range"`
	WeekRange                         string                 `json:"week_range" redis:"week_range"`
	Volume                            int64                  `json:"volume" redis:"volume"`
	AvgVolume                         int64                  `json:"avg_volume" redis:"avg_volume"`
	MarketCap                         float64                `json:"market_cap" redis:"market_cap"`
	Beta                              float64                `json:"beta" redis:"beta"`
	PERatio                           float64                `json:"pe_ratio" redis:"pe_ratio"`
	EPS                               float64                `json:"eps" redis:"eps"`
	DividendYield                     string                 `json:"dividend_yield" redis:"dividend_yield"`
	ExDividendDate                    string                 `json:"ex_dividend_date" redis:"ex_dividend_date"`
	TargetEst                         float64                `json:"target_est" redis:"target_est"`
	URL                               string                 `json:"url" redis:"url"`
	Timestamp                         string                 `json:"timestamp" redis:"timestamp"`
	StockIndex                        string                 `json:"stock_market_index" redis:"stock_market_index"`
	FinancialsCurrency                string                 `json:"fanacials_currency" redis:"fanacials_currency"`
	CompanyProfileAddress             string                 `json:"company_profile_address" redis:"company_profile_address"`
	CompanyProfileHeadquartersCountry string                 `json:"company_profile_headquarters_country" redis:"company_profile_headquarters_country"`
	CompanyProfileCountry             string                 `json:"company_profile_country" redis:"company_profile_country"`
	CompanyProfileWebsite             string                 `json:"company_profile_website" redis:"company_profile_website"`
	CompanyProfilePhone               string                 `json:"company_profile_phone" redis:"company_profile_phone"`
	CompanyProfileSector              string                 `json:"company_profile_sector" redis:"company_profile_sector"`
	CompanyProfileIndustry            string                 `json:"company_profile_industry" redis:"company_profile_industry"`
	CompanyProfileEmployees           int                    `json:"company_profile_employees" redis:"company_profile_employees"`
	CompanyProfileDescription         string                 `json:"company_profile_description" redis:"company_profile_description"`
	ValuationMeasures                 map[string]interface{} `json:"valuation_measures" redis:"valuation_measures"`
	FinancialHighlights               map[string]interface{} `json:"financial_highlights" redis:"financial_highlights"`
	GrowthEstimates                   map[string]interface{} `json:"growth_estimates" redis:"growth_estimates"`

	AnalystPriceTarget          AnalystPriceTarget     `json:"analyst_price_target" redis:"analyst_price_target"`
	CompanyProfileKeyExecutives []CompanyKeyExecutive  `json:"company_profile_key_executives" redis:"company_profile_key_executives"`
	Financials                  []Financial            `json:"financials" redis:"financials"`
	EarningsEstimate            map[string][]ItemValue `json:"earnings_estimate" redis:"earnings_estimate"`
	RevenueEstimate             map[string][]ItemValue `json:"revenue_estimate" redis:"revenue_estimate"`
	EarningsHistory             map[string][]ItemValue `json:"earnings_history" redis:"earnings_history"`
	EPSTrend                    map[string][]ItemValue `json:"eps_trend" redis:"eps_trend"`
	EPSRevisions                map[string][]ItemValue `json:"eps_revisions" redis:"eps_revisions"`
}

func (stock *Stock) IsOk() bool {
	return stock.CompanyID != ""
}

func (s Stock) String() string {
	// Marshal complex map fields
	valMeasures, _ := json.MarshalIndent(s.ValuationMeasures, "    ", "  ")
	finHighlights, _ := json.MarshalIndent(s.FinancialHighlights, "    ", "  ")
	growthEstimates, _ := json.MarshalIndent(s.GrowthEstimates, "    ", "  ")
	keyExecutives, _ := json.MarshalIndent(s.CompanyProfileKeyExecutives, "    ", "  ")
	financials, _ := json.MarshalIndent(s.Financials, "    ", "  ")

	return fmt.Sprintf(
		"Stock: %s (%s)\n"+
			"  Ticker: %s | Exchange: %s | Currency: %s\n"+
			"  Price: %.2f | Open: %.2f | Prev Close: %.2f | Volume: %d\n"+
			"  MarketCap: %.2f | P/E: %.2f | EPS: %.2f | Beta: %.2f\n"+
			"  Dividend Yield: %s | Target Est: %.2f | Earnings Date: %s\n"+
			"  Sector: %s | Industry: %s | Employees: %d\n"+
			"  Website: %s | Phone: %s\n"+
			"  Summary: %s\n"+
			"  Valuation Measures:\n%s\n"+
			"  Financial Highlights:\n%s\n"+
			"  Growth Estimates:\n%s\n"+
			"  Key Executives:\n%s\n"+
			"  Financials:\n%s\n",
		s.Name, s.CompanyID,
		s.StockTicker, s.Exchange, s.Currency,
		s.ClosingPrice, s.Open, s.PreviousClose, s.Volume,
		s.MarketCap, s.PERatio, s.EPS, s.Beta,
		s.DividendYield, s.TargetEst, s.EarningsDate,
		s.CompanyProfileSector, s.CompanyProfileIndustry, s.CompanyProfileEmployees,
		s.CompanyProfileWebsite, s.CompanyProfilePhone,
		s.Summary,
		valMeasures,
		finHighlights,
		growthEstimates,
		keyExecutives,
		financials,
	)
}
