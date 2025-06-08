package models

type Portfolio struct {
	PortfolioId  int64  `json:"portfolio_id" gorm:"primaryKey;autoIncrement"`
	AccountId    int64  `json:"account_id" gorm:"foreignKey:AccountId;references:AccountId;not null"`
	CurrentValue string `json:"current_value" gorm:"not null"`
	ProfitLoss   string `json:"profit_loss" gorm:"not null"`
}
