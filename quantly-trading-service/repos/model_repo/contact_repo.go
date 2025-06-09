package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type ContactRepo struct {
	DB *gorm.DB
}

func (r *ContactRepo) GetById(contactId int64) (*models.Contact, error) {
	contact := &models.Contact{}
	if err := r.DB.First(contact, "contact_id = ?", contactId).Error; err != nil {
		return nil, utils.OnError(err, "can not get contact")
	}
	return contact, nil
}

func (r *ContactRepo) Create(contact *models.Contact) error {
	err := r.DB.Create(contact).Error
	return utils.OnError(err, "can not create contact")
}

func (r *ContactRepo) Update(contact *models.Contact) error {
	if contact.ContactId == 0 {
		return gorm.ErrRecordNotFound
	}

	err := r.DB.Save(contact).Error
	return utils.OnError(err, "can not update contact")

}

func (r *ContactRepo) DeleteById(contactId int64) error {
	err := r.DB.Where("contact_id = ?", contactId).Delete(&models.Contact{}).Error
	return utils.OnError(err, "can not delete contact")
}
