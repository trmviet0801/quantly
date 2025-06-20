package models

type Portfolio struct {
	PortfolioId  string `json:"portfolio_id" gorm:"primaryKey;type:varchar(64)"`
	AccountId    string `json:"account_id" gorm:"foreignKey:AccountId;references:AccountId;not null;type:varchar(64)"`
	CurrentValue string `json:"current_value" gorm:"not null"`
	ProfitLoss   string `json:"profit_loss" gorm:"not null"`
}
