package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type DisclosureRepo struct {
	DB *gorm.DB
}

func (r *DisclosureRepo) GetById(disclosureId int64) (*models.Disclosure, error) {
	disclosure := &models.Disclosure{}
	if err := r.DB.First(disclosure, "disclosure_id = ?", disclosureId).Error; err != nil {
		return nil, utils.OnError(err, "can not get disclosure")
	}
	return disclosure, nil
}

func (r *DisclosureRepo) Create(disclosure *models.Disclosure) error {
	err := r.DB.Create(disclosure).Error
	return utils.OnError(err, "can not create disclosure")
}

func (r *DisclosureRepo) Update(disclosure *models.Disclosure) error {
	if disclosure.DisclosureId == 0 {
		return fmt.Errorf("can not update disclosure: invalid input")
	}

	err := r.DB.Save(disclosure).Error
	return utils.OnError(err, "can not update disclosure")
}

func (r *DisclosureRepo) DeleteById(disclosureId int64) error {
	err := r.DB.Where("disclosureId = ?", disclosureId).Delete(&models.Disclosure{}).Error
	return utils.OnError(err, "can not delete disclosure")
}
