package models

type TradingConfiguration struct {
	TradingConfigurationId int64  `json:"trading_configuration_id" binding:"required" gorm:"primaryKey"`
	AccountId              string `json:"account_id" binding:"required" gorm:"foreignKey:AccountId;references:AccountId;type:varchar(64)"`
	DtbpCheck              string `json:"dtbp_check" binding:"required"`            // "true" or "false"
	NoShhorting            bool   `json:"no_shorting" binding:"required"`           // "true" or "false"
	FractionalTrading      bool   `json:"fractional_trading" binding:"required"`    // "true" or "false"
	MaxMarginMultiplier    string `json:"max_margin_multiplier" binding:"required"` // "0.5", "1.0", etc.
}
