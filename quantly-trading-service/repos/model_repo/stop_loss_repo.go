package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type StopLossRepo struct {
	db *gorm.DB
}

func (r *StopLossRepo) GetById(stopLossId string) (*models.StopLoss, error) {
	var stopLoss *models.StopLoss
	err := r.db.First(stopLoss, "stop_loss_id = ?", stopLossId).Error
	if err != nil {
		return nil, fmt.Errorf("can not get stop loss by id: %w", err)
	}
	return stopLoss, nil
}

func (r *StopLossRepo) Create(stopLoss *models.StopLoss) error {
	err := r.db.Create(stopLoss).Error
	if err != nil {
		return fmt.Errorf("can not create stop loss: %w", err)
	}
	return nil
}

func (r *StopLossRepo) DeleteById(stopLossId string) error {
	err := r.db.Where("stop_loss_id = ?", stopLossId).Delete(&models.StopLoss{}).Error
	if err != nil {
		return fmt.Errorf("can not delete stop loss: %w", err)
	}
	return nil
}

func (r *StopLossRepo) Update(stopLoss *models.StopLoss) error {
	if stopLoss.StopLostId == 0 {
		return fmt.Errorf("can not update stop loss: input invalid")
	}

	err := r.db.Save(stopLoss).Error
	if err != nil {
		return fmt.Errorf("can not update stop loss")
	}
	return nil
}
