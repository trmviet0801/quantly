package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type TrustedContactRepo struct {
	db *gorm.DB
}

func (r *TrustedContactRepo) GetById(trustedContactId int64) (*models.TrustedContact, error) {
	var trustedContact *models.TrustedContact
	if err := r.db.First(trustedContact, "trusted_contact_id = ?", trustedContactId).Error; err != nil {
		return nil, utils.OnError(err, "can not select trusted contact")
	}
	return trustedContact, nil
}

func (r *TrustedContactRepo) Create(trustedContact *models.TrustedContact) error {
	err := r.db.Create(trustedContact).Error
	return utils.OnError(err, "can not create trusted contact")
}

func (r *TrustedContactRepo) Update(trustedContact *models.TrustedContact) error {
	if trustedContact.TrustedContactId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(trustedContact).Error
	return utils.OnError(err, "can not update trusted contact")
}

func (r *TrustedContactRepo) DeleteById(trustedContactId int64) error {
	err := r.db.Where("trusted_contact_id = ?", trustedContactId).Delete(&models.TrustedContact{}).Error
	return utils.OnError(err, "can not delete trusted contact")
}
