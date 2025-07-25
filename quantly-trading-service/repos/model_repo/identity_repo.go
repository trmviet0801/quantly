package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type IdentityRepo struct {
	DB *gorm.DB
}

func (r *IdentityRepo) GetById(identityId int64) (*models.Identity, error) {
	identity := &models.Identity{}
	if err := r.DB.First(identity, "identity_id = ?", identityId).Error; err != nil {
		return nil, utils.OnError(err, "can not get identity")
	}
	return identity, nil
}

func (r *IdentityRepo) Create(identity *models.Identity) error {
	err := r.DB.Create(identity).Error
	return utils.OnError(err, "can not create identity")
}

func (r *IdentityRepo) Update(identity *models.Identity) error {
	if identity.IdentityId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.DB.Save(identity).Error
	return utils.OnError(err, "can not update identity")
}

func (r *IdentityRepo) DeleteById(identityId int64) error {
	err := r.DB.Where("identity_id = ?", identityId).Delete(&models.Identity{}).Error
	return utils.OnError(err, "can not delete identity")
}
