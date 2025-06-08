package models

import "time"

type KycResult struct {
	KycResultId    int64     `json:"kyc_result_id" gorm:"primaryKey"`
	AccountId      int64     `json:"account_id" gorm:"foreignKey:AccountId;references:AccountId;not null"`
	Status         string    `json:"status"`
	InternalStatus string    `json:"internal_status"`
	TimeStamp      time.Time `json:"timestamp"`
}
