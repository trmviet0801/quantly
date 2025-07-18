package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type TradeSignalRepo struct {
	DB *gorm.DB
}

func (r *TradeSignalRepo) GetById(tradeSignalId int64) (*models.TradeSignal, error) {
	tradeSignal := &models.TradeSignal{}
	if err := r.DB.First(tradeSignal, "trade_signal_id = ?", tradeSignalId).Error; err != nil {
		return nil, utils.OnError(err, "can not get trade signal")
	}

	return tradeSignal, nil
}

func (r *TradeSignalRepo) Create(tradeSignal *models.TradeSignal) error {
	err := r.DB.Create(tradeSignal).Error
	return utils.OnError(err, "can not create trade signal")
}

func (r *TradeSignalRepo) Update(tradeSignal *models.TradeSignal) error {
	if tradeSignal.TradeSignalId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.DB.Save(tradeSignal).Error
	return utils.OnError(err, "can not update trade signal")
}

func (r *TradeSignalRepo) DeleteById(tradeSignalId int64) error {
	err := r.DB.Where("trade_signal_id = ?", tradeSignalId).Delete(&models.TradeSignal{}).Error
	return utils.OnError(err, "can not delete trade signal")
}
