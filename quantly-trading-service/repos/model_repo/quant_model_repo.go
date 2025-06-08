package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type QuantModelRepo struct {
	db *gorm.DB
}

func (r *QuantModelRepo) GetById(quantModelId int64) (*models.QuantModel, error) {
	var quantModel *models.QuantModel
	err := r.db.First(quantModel, "quant_model_id = ?", quantModelId).Error
	if err != nil {
		return nil, err
	}
	return quantModel, nil
}

func (r *QuantModelRepo) Create(quantModel *models.QuantModel) error {
	err := r.db.Create(quantModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *QuantModelRepo) Update(quantModel *models.QuantModel) error {
	if quantModel.QuantModelId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(quantModel).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *QuantModelRepo) DeleteById(quantModelId int64) error {
	err := r.db.Where("quant_model_id = ?", quantModelId).Delete(&models.QuantModel{}).Error
	if err != nil {
		return err
	}
	return nil
}
