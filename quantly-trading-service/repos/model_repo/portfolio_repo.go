package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type PortfolioRepo struct {
	db *gorm.DB
}

func (r *PortfolioRepo) GetById(portfolioId int64) (*models.Portfolio, error) {
	var portfolio *models.Portfolio
	err := r.db.First(portfolio, "portfolio_id = ?", portfolioId).Error
	if err != nil {
		return nil, err
	}
	return portfolio, nil
}

func (r *PortfolioRepo) Create(portfolio *models.Portfolio) error {
	err := r.db.Create(portfolio).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PortfolioRepo) Update(portfolio *models.Portfolio) error {
	if portfolio.PortfolioId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(portfolio).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PortfolioRepo) DeleteById(portfolioId int64) error {
	err := r.db.Where("portfolio_id = ?", portfolioId).Delete(&models.Portfolio{}).Error
	if err != nil {
		return err
	}
	return nil
}
