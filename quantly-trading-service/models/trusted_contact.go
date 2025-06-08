package models

type TrustedContact struct {
	TrustedContactId int64  `json:"trusted_contact_id" gorm:"primaryKey;autoIncrement"`
	AccountId        int64  `json:"account_id" gorm:"foreignKey:AccountId;references:AccountId;not null"`
	GivenName        string `json:"given_name" gorm:"not null"`
	FamilyName       string `json:"family_name"`
	EmailAddress     string `json:"email_address" gorm:"not null"`
}
