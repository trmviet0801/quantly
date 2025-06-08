package models

import "time"

type Identity struct {
	IdentityId            int64     `json:"identity_id" gorm:"primaryKey"`
	AccountId             int64     `json:"account_id" gorm:"foreignKey:AccountId;references:AccountId;not null"`
	TaxId                 string    `json:"tax_id"`
	TaxIdType             string    `json:"tax_id_type"` // e.g., "SSN", "EIN", etc.
	CountryOfCitizenship  string    `json:"country_of_citizenship"`
	CountryOfBirth        string    `json:"country_of_birth"`
	CountryOfTaxResidence string    `json:"country_of_tax_residence"`
	FundingSource         string    `json:"funding_source"` // JSON object to store funding source details
	DateOfBirth           time.Time `json:"date_of_birth"`
}
