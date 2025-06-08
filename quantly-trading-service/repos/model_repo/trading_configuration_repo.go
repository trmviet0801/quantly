package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type TradingConfigurationRepo struct {
	db *gorm.DB
}

func (r *TradingConfigurationRepo) GetById(tradingConfigurationId int64) (*models.TradingConfiguration, error) {
	var tradingConfiguration *models.TradingConfiguration
	err := r.db.First(tradingConfiguration, "trading_configuration_id = ?", tradingConfigurationId).Error
	if err != nil {
		return nil, err
	}
	return tradingConfiguration, nil
}

func (r *TradingConfigurationRepo) Create(tradingConfiguration *models.TradingConfiguration) error {
	err := r.db.Create(tradingConfiguration).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TradingConfigurationRepo) Update(tradingConfiguration *models.TradingConfiguration) error {
	if tradingConfiguration.TradingConfigurationId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(tradingConfiguration).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TradingConfigurationRepo) DeleteById(tradingConfigurationId int64) error {
	err := r.db.Where("trading_configuration_id = ?", tradingConfigurationId).Delete(&models.TradingConfiguration{}).Error
	if err != nil {
		return err
	}
	return nil
}
