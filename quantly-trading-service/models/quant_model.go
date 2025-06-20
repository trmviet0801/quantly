package models

type QuantModel struct {
	QuantModelId string `json:"quant_model_id" gorm:"primaryKey;type:varchar(64)"`
	Name         string `json:"name" gorm:"not null"`
}
