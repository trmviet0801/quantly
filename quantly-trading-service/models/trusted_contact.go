package models

type TrustedContact struct {
	TrustedContactId int64  `json:"trusted_contact_id" gorm:"primaryKey;autoIncrement"`
	AccountId        string `json:"account_id" gorm:"foreignKey:AccountId;references:AccountId;not null;type:varchar(64)"`
	GivenName        string `json:"given_name" gorm:"not null"`
	FamilyName       string `json:"family_name"`
	EmailAddress     string `json:"email_address" gorm:"not null"`
}
