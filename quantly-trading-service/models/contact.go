package models

type Contact struct {
	ContactId      int64  `json:"contact_id" gorm:"primaryKey;autoIncrement"`
	AccountId      int64  `json:"account_id" gorm:"foreignKey:AccountId;references:AccountId;not null"`
	EmailAddress   string `json:"email_address" gorm:"not null"`
	PhoneNumber    string `json:"phone_number" gorm:"not null"`
	StreestAddress string `json:"street_address"`
	City           string `json:"city"`
	State          string `json:"state"`
	PostalCode     string `json:"postal_code"`
	Country        string `json:"country"`
	GivenName      string `json:"given_name"`
	MiddleName     string `json:"middle_name"`
	FamilyName     string `json:"family_name"`
}
