package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type ContactRepo struct {
	db *gorm.DB
}

func (r *ContactRepo) GetById(contactId int64) (*models.Contact, error) {
	contact := &models.Contact{}
	if err := r.db.First(contact, "contact_id = ?", contactId).Error; err != nil {
		return nil, utils.OnError(err, "can not get contact")
	}
	return contact, nil
}

func (r *ContactRepo) Create(contact *models.Contact) error {
	err := r.db.Create(contact).Error
	return utils.OnError(err, "can not create contact")
}

func (r *ContactRepo) Update(contact *models.Contact) error {
	if contact.ContactId == 0 {
		return gorm.ErrRecordNotFound
	}

	err := r.db.Save(contact).Error
	return utils.OnError(err, "can not update contact")

}

func (r *ContactRepo) DeleteById(contactId int64) error {
	err := r.db.Where("contact_id = ?", contactId).Delete(&models.Contact{}).Error
	return utils.OnError(err, "can not delete contact")
}
