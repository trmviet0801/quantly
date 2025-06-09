package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type StopLossRepo struct {
	DB *gorm.DB
}

func (r *StopLossRepo) GetById(stopLossId string) (*models.StopLoss, error) {
	stopLoss := &models.StopLoss{}
	if err := r.DB.First(stopLoss, "stop_loss_id = ?", stopLossId).Error; err != nil {
		return nil, utils.OnError(err, "can not get stop loss")
	}
	return stopLoss, nil
}

func (r *StopLossRepo) Create(stopLoss *models.StopLoss) error {
	err := r.DB.Create(stopLoss).Error
	return utils.OnError(err, "can not create stop loss")
}

func (r *StopLossRepo) DeleteById(stopLossId string) error {
	err := r.DB.Where("stop_loss_id = ?", stopLossId).Delete(&models.StopLoss{}).Error
	return utils.OnError(err, "can not delete stop loss")
}

func (r *StopLossRepo) Update(stopLoss *models.StopLoss) error {
	if stopLoss.StopLostId == 0 {
		return fmt.Errorf("can not update stop loss: input invalid")
	}

	err := r.DB.Save(stopLoss).Error
	return utils.OnError(err, "can not update stop loss")
}
