package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type QuantModelRepo struct {
	db *gorm.DB
}

func (r *QuantModelRepo) GetById(quantModelId int64) (*models.QuantModel, error) {
	var quantModel *models.QuantModel
	if err := r.db.First(quantModel, "quant_model_id = ?", quantModelId).Error; err != nil {
		return nil, utils.OnError(err, "can not get quant model")
	}
	return quantModel, nil
}

func (r *QuantModelRepo) Create(quantModel *models.QuantModel) error {
	err := r.db.Create(quantModel).Error
	return utils.OnError(err, "can not create quant model")
}

func (r *QuantModelRepo) Update(quantModel *models.QuantModel) error {
	if quantModel.QuantModelId == 0 {
		return gorm.ErrRecordNotFound
	}

	err := r.db.Save(quantModel).Error
	return utils.OnError(err, "can not update quant model")
}
func (r *QuantModelRepo) DeleteById(quantModelId int64) error {
	err := r.db.Where("quant_model_id = ?", quantModelId).Delete(&models.QuantModel{}).Error
	return utils.OnError(err, "can not delete quant model")
}
