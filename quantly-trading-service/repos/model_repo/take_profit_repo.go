package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type TakeProfitRepo struct {
	db *gorm.DB
}

func (r *TakeProfitRepo) GetById(takeProfitId string) (*models.TakeProfit, error) {
	var takeProfit *models.TakeProfit
	err := r.db.First(takeProfit, "take_profit_id = ?", takeProfitId).Error
	if err != nil {
		return nil, fmt.Errorf("can not get take profit by id: %w", err)
	}
	return takeProfit, nil
}
func (r *TakeProfitRepo) Create(takeProfit *models.TakeProfit) error {
	err := r.db.Create(takeProfit).Error
	if err != nil {
		return fmt.Errorf("can not create take profit: %w", err)
	}
	return nil
}
func (r *TakeProfitRepo) Update(takeProfit *models.TakeProfit) error {
	if takeProfit.TakeProfitId == 0 {
		return fmt.Errorf("can not update take profit: invalid input")
	}
	err := r.db.Save(takeProfit).Error
	if err != nil {
		return fmt.Errorf("can not update take profit: %w", err)
	}
	return nil
}
func (r *TakeProfitRepo) DeleteById(takeProfitId string) error {
	err := r.db.Where("take_profit_id = ?", takeProfitId).Delete(&models.TakeProfit{}).Error
	if err != nil {
		return fmt.Errorf("can not delete take profit: %w", err)
	}
	return nil
}
