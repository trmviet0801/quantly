package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type ContactRepo struct {
	db *gorm.DB
}

func (r *ContactRepo) GetById(contactId int64) (*models.Contact, error) {
	var contact *models.Contact
	err := r.db.First(contact, "contact_id = ?", contactId).Error
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (r *ContactRepo) Create(contact *models.Contact) error {
	err := r.db.Create(contact).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepo) Update(contact *models.Contact) error {
	if contact.ContactId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(contact).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepo) DeleteById(contactId int64) error {
	err := r.db.Where("contact_id = ?", contactId).Delete(&models.Contact{}).Error
	if err != nil {
		return err
	}
	return nil
}
