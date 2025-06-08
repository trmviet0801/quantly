package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type TradeSignalRepo struct {
	db *gorm.DB
}

func (r *TradeSignalRepo) GetById(tradeSignalId int64) (*models.TradeSignal, error) {
	var tradeSignal *models.TradeSignal
	err := r.db.First(tradeSignal, "trade_signal_id = ?", tradeSignalId).Error
	if err != nil {
		return nil, err
	}
	return tradeSignal, nil
}

func (r *TradeSignalRepo) Create(tradeSignal *models.TradeSignal) error {
	err := r.db.Create(tradeSignal).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TradeSignalRepo) Update(tradeSignal *models.TradeSignal) error {
	if tradeSignal.TradeSignalId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(tradeSignal).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TradeSignalRepo) DeleteById(tradeSignalId int64) error {
	err := r.db.Where("trade_signal_id = ?", tradeSignalId).Delete(&models.TradeSignal{}).Error
	if err != nil {
		return err
	}
	return nil
}
