package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type PortfolioRepo struct {
	db *gorm.DB
}

func (r *PortfolioRepo) GetById(portfolioId int64) (*models.Portfolio, error) {
	portfolio := &models.Portfolio{}
	if err := r.db.First(portfolio, "portfolio_id = ?", portfolioId).Error; err != nil {
		return nil, utils.OnError(err, "can not get portfolio")
	}
	return portfolio, nil
}

func (r *PortfolioRepo) Create(portfolio *models.Portfolio) error {
	err := r.db.Create(portfolio).Error
	return utils.OnError(err, "can not create portfolio")
}

func (r *PortfolioRepo) Update(portfolio *models.Portfolio) error {
	if portfolio.PortfolioId == 0 {
		return gorm.ErrRecordNotFound
	}

	err := r.db.Save(portfolio).Error
	return utils.OnError(err, "can not update portfolio")
}

func (r *PortfolioRepo) DeleteById(portfolioId int64) error {
	err := r.db.Where("portfolio_id = ?", portfolioId).Delete(&models.Portfolio{}).Error
	return utils.OnError(err, "can not delete profolio")
}
