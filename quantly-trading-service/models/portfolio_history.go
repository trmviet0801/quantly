package models

import (
	"fmt"
	"strings"

	customtype "github.com/trmviet0801/quantly/models/custom_type"
)

type PortfolioHistory struct {
	AccountId     string                  `json:"account_id" gorm:"primaryKey;type:varchar(64)"`
	Timestamp     customtype.Int64Slice   `json:"timestamp" gorm:"type:json"`
	Equity        customtype.Float64Slice `json:"equity" gorm:"type:json"`
	ProfitLoss    customtype.Float64Slice `json:"profit_loss" gorm:"type:json"`
	ProfitLossPct customtype.Float64Slice `json:"profit_loss_pct" gorm:"type:json"`
	BaseValue     float64                 `json:"base_value"`
	BaseValueAsof string                  `json:"base_value_asof"`
	Timeframe     string                  `json:"timeframe"`
}

func (p PortfolioHistory) String() string {
	return fmt.Sprintf(
		`PortfolioHistory:
  AccountId: %s
  Timeframe: %s
  BaseValue: %.2f
  BaseValueAsof: %s
  Timestamp: %s
  Equity: %s
  ProfitLoss: %s
  ProfitLossPct: %s`,
		p.AccountId,
		p.Timeframe,
		p.BaseValue,
		p.BaseValueAsof,
		int64SliceToString(p.Timestamp),
		float64SliceToString(p.Equity),
		float64SliceToString(p.ProfitLoss),
		float64SliceToString(p.ProfitLossPct),
	)
}

func int64SliceToString(slice []int64) string {
	strs := make([]string, len(slice))
	for i, v := range slice {
		strs[i] = fmt.Sprintf("%d", v)
	}
	return "[" + strings.Join(strs, ", ") + "]"
}

func float64SliceToString(slice []float64) string {
	strs := make([]string, len(slice))
	for i, v := range slice {
		strs[i] = fmt.Sprintf("%.2f", v)
	}
	return "[" + strings.Join(strs, ", ") + "]"
}
