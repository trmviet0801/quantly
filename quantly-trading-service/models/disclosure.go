package models

type Disclosure struct {
	DisclosureId                int64 `json:"disclosure_id" gorm:"primarKey"`
	AccountId                   int64 `json:"account_id" gorm:"foreignKey:AccountId;references:AccountId;not null"`
	IsControlPerson             bool  `json:"is_control_person"`
	IsAffiliatedExchangeOrFinra bool  `json:"is_affiliated_exchanged_or_finra"`
	ImmediateFamilyExposed      bool  `json:"immediate_family_exposed"`
}
