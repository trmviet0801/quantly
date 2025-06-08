package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type KycResultRepo struct {
	db *gorm.DB
}

func (r *KycResultRepo) GetById(kycResultId int) (*models.KycResult, error) {
	var kycResult *models.KycResult
	err := r.db.First(kycResult, "kyc_result_id = ?", kycResultId).Error
	if err != nil {
		return nil, fmt.Errorf("can not get kyc result: %w", err)
	}
	return kycResult, nil
}

func (r *KycResultRepo) Create(kycResult *models.KycResult) error {
	err := r.db.Create(kycResult).Error
	if err != nil {
		return fmt.Errorf("can not create kyc result: %w", err)
	}
	return nil
}

func (r *KycResultRepo) Update(kycResult *models.KycResult) error {
	if kycResult.KycResultId == 0 {
		return fmt.Errorf("can not update kyc result: input invalid")
	}

	err := r.db.Save(kycResult).Error
	if err != nil {
		return fmt.Errorf("can not update kyc: %w", err)
	}
	return nil
}

func (r *KycResultRepo) DeleteById(kycResultId int64) error {
	err := r.db.Where("kyc_result_id = ?", kycResultId).Delete(&models.KycResult{}).Error
	if err != nil {
		return fmt.Errorf("can not delete kyc result: %w", err)
	}
	return nil
}
