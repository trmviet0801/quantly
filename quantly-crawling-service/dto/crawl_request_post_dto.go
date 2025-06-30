package dto

import (
	"fmt"
	"os"
)

type CrawlRequestPostDto struct {
	Input              []InputURL `json:"input"`
	CustomOutputFields []string   `json:"custom_output_fields"`
}

func (c *CrawlRequestPostDto) ConstructBody(stockIds []string) {
	c.SetUpDefaultCustomFields()
	c.ConstructInputUrl(stockIds)
}

// Add URL for crawling
func (c *CrawlRequestPostDto) ConstructInputUrl(stockIds []string) {
	for _, stockId := range stockIds {
		c.Input = append(c.Input, InputURL{URL: fmt.Sprintf("%s%s", os.Getenv("FINANCE_YAHOO_URL"), stockId)})
	}
}

// Add pre-definded fields that will be crawled by BrightData on Finance.yahoo
func (c *CrawlRequestPostDto) SetUpDefaultCustomFields() {
	c.CustomOutputFields = []string{
		"name",
		"company_id",
		"entity_type",
		"summary",
		"stock_ticker",
		"currency",
		"earnings_date",
		"exchange",
		"closing_price",
		"previous_close",
		"open",
		"bid",
		"ask",
		"day_range",
		"week_range",
		"volume",
		"avg_volume",
		"market_cap",
		"beta",
		"pe_ratio",
		"eps",
		"dividend_yield",
		"ex_dividend_date",
		"target_est",
		"url",
		"people_also_watch",
		"similar",
		"risk_score",
		"risk_score_text",
		"risk_score_percentile",
		"recommendation_rating",
		"analyst_price_target",
		"company_profile_address",
		"company_profile_headquarters_country",
		"company_profile_country",
		"company_profile_website",
		"company_profile_phone",
		"company_profile_sector",
		"company_profile_industry",
		"company_profile_employees",
		"company_profile_key_executives",
		"company_profile_description",
		"valuation_measures",
		"financial_highlights",
		"financials",
		"financials_quarterly",
		"earnings_estimate",
		"revenue_estimate",
		"earnings_history",
		"eps_trend",
		"eps_revisions",
		"growth_estimates",
		"top_analysts",
		"upgrades_and_downgrades",
		"recent_news",
		"fanacials_currency",
		"stock_market_index",
		"timestamp",
		"input",
		"error",
		"error_code",
		"warning",
		"warning_code",
	}
}
