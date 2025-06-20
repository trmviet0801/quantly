package models

import (
	"encoding/json"
	"time"
)

type Account struct {
	AccountId              string          `json:"account_id" gorm:"primaryKey;type:varchar(64)"`
	UserId                 string          `binding:"required" json:"user_id" gorm:"type:varchar(64)"`
	AccountNumber          string          `json:"account_number"`
	Status                 string          `json:"status"`
	CryptoStatus           string          `json:"crypto_status"`
	Currency               string          `json:"currency"`
	LastEquity             string          `json:"last_equity"`
	CreatedAt              time.Time       `json:"created_at"`
	AccountType            string          `json:"account_type"`
	EnabledAssets          json.RawMessage `json:"enabled_assets"` // JSON array of enabled assets
	PartnerUserId          string          `json:"partner_user_id"`
	FundingInstructionsUrl string          `json:"funding_instructions_url"`
	PatternDayTrader       int8            `json:"pattern_day_trader"` // 0 or 1
	KycCompletedAt         time.Time       `json:"kyc_completed_at"`
	KycStatus              string          `json:"kyc_status"`
	AccountAtype           string          `json:"account_atype"`
	ManagementStatus       string          `json:"management_status"`
	ClearingBroker         string          `json:"clearing_broker"`
	ClearingAccountNumber  string          `json:"clearing_account_number"`
	Contacts               []Contact       `gorm:"foreignKey:AccountId;references:AccountId"`
}
