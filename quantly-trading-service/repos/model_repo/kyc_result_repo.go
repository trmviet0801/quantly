package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type KycResultRepo struct {
	db *gorm.DB
}

func (r *KycResultRepo) GetById(kycResultId int) (*models.KycResult, error) {
	var kycResult *models.KycResult
	if err := r.db.First(kycResult, "kyc_result_id = ?", kycResultId).Error; err != nil {
		return nil, utils.OnError(err, "can not get kyc result")
	}
	return kycResult, nil
}

func (r *KycResultRepo) Create(kycResult *models.KycResult) error {
	err := r.db.Create(kycResult).Error
	return utils.OnError(err, "can not create kyc result")
}

func (r *KycResultRepo) Update(kycResult *models.KycResult) error {
	if kycResult.KycResultId == 0 {
		return fmt.Errorf("can not update kyc result: input invalid")
	}

	err := r.db.Save(kycResult).Error
	return utils.OnError(err, "can not update kyc")
}

func (r *KycResultRepo) DeleteById(kycResultId int64) error {
	err := r.db.Where("kyc_result_id = ?", kycResultId).Delete(&models.KycResult{}).Error
	return utils.OnError(err, "can not delete kyc result")
}
