package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (r *UserRepo) GetById(userId int64) (*models.User, error) {
	var user *models.User
	err := r.db.First(user, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) Create(user *models.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Update(user *models.User) error {
	if user.UserId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) DeleteById(userId int64) error {
	err := r.db.Where("user_id = ?", userId).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
