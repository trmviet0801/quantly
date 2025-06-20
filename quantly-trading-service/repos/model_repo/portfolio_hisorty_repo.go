package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type PortfolioHistoryRepo struct {
	DB *gorm.DB
}

// GetByAccountId retrieves a PortfolioHistory by AccountId (primary key)
func (r *PortfolioHistoryRepo) GetByAccountId(accountId string) (*models.PortfolioHistory, error) {
	history := &models.PortfolioHistory{}
	err := r.DB.First(history, "account_id = ?", accountId).Error
	if err != nil {
		return nil, utils.OnError(err, "can not get portfolio history")
	}
	return history, nil
}

// Create inserts a new PortfolioHistory record
func (r *PortfolioHistoryRepo) Create(history *models.PortfolioHistory) error {
	err := r.DB.Create(history).Error
	return utils.OnError(err, "can not create portfolio history")
}

// Update updates an existing PortfolioHistory
func (r *PortfolioHistoryRepo) Update(history *models.PortfolioHistory) error {
	if history.AccountId == "" {
		return fmt.Errorf("can not update portfolio history: invalid input")
	}
	err := r.DB.Save(history).Error
	return utils.OnError(err, "can not update portfolio history")
}

// DeleteByAccountId deletes a PortfolioHistory record by AccountId
func (r *PortfolioHistoryRepo) DeleteByAccountId(accountId string) error {
	err := r.DB.Where("account_id = ?", accountId).Delete(&models.PortfolioHistory{}).Error
	return utils.OnError(err, "can not delete portfolio history")
}
