package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type TrustedContactRepo struct {
	db *gorm.DB
}

func (r *TrustedContactRepo) GetById(trustedContactId int64) (*models.TrustedContact, error) {
	var trustedContact *models.TrustedContact
	err := r.db.First(trustedContact, "trusted_contact_id = ?", trustedContactId).Error
	if err != nil {
		return nil, err
	}
	return trustedContact, nil
}

func (r *TrustedContactRepo) Create(trustedContact *models.TrustedContact) error {
	err := r.db.Create(trustedContact).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TrustedContactRepo) Update(trustedContact *models.TrustedContact) error {
	if trustedContact.TrustedContactId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(trustedContact).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TrustedContactRepo) DeleteById(trustedContactId int64) error {
	err := r.db.Where("trusted_contact_id = ?", trustedContactId).Delete(&models.TrustedContact{}).Error
	if err != nil {
		return err
	}
	return nil
}
