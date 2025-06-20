package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type TakeProfitRepo struct {
	DB *gorm.DB
}

func (r *TakeProfitRepo) GetById(takeProfitId string) (*models.TakeProfit, error) {
	takeProfit := &models.TakeProfit{}
	if err := r.DB.First(takeProfit, "take_profit_id = ?", takeProfitId).Error; err != nil {
		return nil, utils.OnError(err, "can not get take profit by id")
	}
	return takeProfit, nil
}
func (r *TakeProfitRepo) Create(takeProfit *models.TakeProfit) error {
	err := r.DB.Create(takeProfit).Error
	return utils.OnError(err, "can not create take profit")
}
func (r *TakeProfitRepo) Update(takeProfit *models.TakeProfit) error {
	if takeProfit.TakeProfitId == "" {
		return fmt.Errorf("can not update take profit: invalid input")
	}

	err := r.DB.Save(takeProfit).Error
	return utils.OnError(err, "can not update take profit")
}
func (r *TakeProfitRepo) DeleteById(takeProfitId string) error {
	err := r.DB.Where("take_profit_id = ?", takeProfitId).Delete(&models.TakeProfit{}).Error
	return utils.OnError(err, "can not delete take profit")
}
