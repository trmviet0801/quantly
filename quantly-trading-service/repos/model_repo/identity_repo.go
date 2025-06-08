package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type IdentityRepo struct {
	db *gorm.DB
}

func (r *IdentityRepo) GetById(identityId int64) (*models.Identity, error) {
	var identity *models.Identity
	err := r.db.First(identity, "identity_id = ?", identityId).Error
	if err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *IdentityRepo) Create(identity *models.Identity) error {
	err := r.db.Create(identity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *IdentityRepo) Update(identity *models.Identity) error {
	if identity.IdentityId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(identity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *IdentityRepo) DeleteById(identityId int64) error {
	err := r.db.Where("identity_id = ?", identityId).Delete(&models.Identity{}).Error
	if err != nil {
		return err
	}
	return nil
}
