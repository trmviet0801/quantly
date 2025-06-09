package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type TradingConfigurationRepo struct {
	db *gorm.DB
}

func (r *TradingConfigurationRepo) GetById(tradingConfigurationId int64) (*models.TradingConfiguration, error) {
	tradingConfiguration := &models.TradingConfiguration{}
	if err := r.db.First(tradingConfiguration, "trading_configuration_id = ?", tradingConfigurationId).Error; err != nil {
		return nil, utils.OnError(err, "can not select trading confiuration")
	}
	return tradingConfiguration, nil
}

func (r *TradingConfigurationRepo) Create(tradingConfiguration *models.TradingConfiguration) error {
	err := r.db.Create(tradingConfiguration).Error
	return utils.OnError(err, "can not create trading")
}

func (r *TradingConfigurationRepo) Update(tradingConfiguration *models.TradingConfiguration) error {
	if tradingConfiguration.TradingConfigurationId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(tradingConfiguration).Error
	return utils.OnError(err, "can not update trade signal")
}

func (r *TradingConfigurationRepo) DeleteById(tradingConfigurationId int64) error {
	err := r.db.Where("trading_configuration_id = ?", tradingConfigurationId).Delete(&models.TradingConfiguration{}).Error
	return utils.OnError(err, "can not delete trading configuration")
}
