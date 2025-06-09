package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (r *UserRepo) GetById(userId int64) (*models.User, error) {
	user := &models.User{}
	if err := r.db.First(user, "user_id = ?", userId).Error; err != nil {
		return nil, utils.OnError(err, "can not get user")
	}

	return user, nil
}

func (r *UserRepo) Create(user *models.User) error {
	err := r.db.Create(user).Error
	return utils.OnError(err, "can not create user")
}

func (r *UserRepo) Update(user *models.User) error {
	if user.UserId == 0 {
		return gorm.ErrRecordNotFound
	}

	err := r.db.Save(user).Error
	return utils.OnError(err, "can not update user")
}

func (r *UserRepo) DeleteById(userId int64) error {
	err := r.db.Where("user_id = ?", userId).Delete(&models.User{}).Error
	return utils.OnError(err, "can not delete user")
}
