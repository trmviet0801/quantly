package models

type QuantModel struct {
	QuantModelId int64  `json:"quant_model_id" gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name" gorm:"not null"`
}
