package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func (accountRepo *AccountRepo) GetById(id string) (*models.Account, error) {
	var account models.Account
	if err := accountRepo.db.First(&account, "account_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (accountRepo *AccountRepo) Create(account *models.Account) error {
	err := accountRepo.db.Create(account).Error
	return utils.OnError(err, "can not create new account")
}

func (r *AccountRepo) Update(account *models.Account) error {
	if account.AccountId == 0 {
		return fmt.Errorf("input not valid")
	}
	err := r.db.Save(account).Error
	return utils.OnError(err, "can not update account")
}

func (r *AccountRepo) DeleteById(id string) error {
	err := r.db.Where("account_id = ?", id).Delete(&models.Account{}).Error
	return utils.OnError(err, "can not delete account")
}
