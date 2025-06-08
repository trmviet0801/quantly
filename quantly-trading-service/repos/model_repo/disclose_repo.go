package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type DisclosureRepo struct {
	db *gorm.DB
}

func (r *DisclosureRepo) GetById(disclosureId int64) (*models.Disclosure, error) {
	var disclosure *models.Disclosure
	err := r.db.First(disclosure, "disclosure_id = ?", disclosureId).Error
	if err != nil {
		return nil, fmt.Errorf("can not get disclosure: %w", err)
	}
	return disclosure, nil
}

func (r *DisclosureRepo) Create(disclosure *models.Disclosure) error {
	err := r.db.Create(disclosure).Error
	if err != nil {
		return fmt.Errorf("can not create disclosure: %w", err)
	}
	return nil
}

func (r *DisclosureRepo) Update(disclosure *models.Disclosure) error {
	if disclosure.DisclosureId == 0 {
		return fmt.Errorf("can not update disclosure: invalid input")
	}

	err := r.db.Save(disclosure).Error
	if err != nil {
		return fmt.Errorf("can not update disclosure: %w", err)
	}
	return nil
}

func (r *DisclosureRepo) DeleteById(disclosureId int64) error {
	err := r.db.Where("disclosureId = ?", disclosureId).Delete(&models.Disclosure{}).Error
	if err != nil {
		return fmt.Errorf("can not delete disclosure: %w", err)
	}
	return nil
}
