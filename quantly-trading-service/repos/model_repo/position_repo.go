package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type PositionRepo struct {
	db *gorm.DB
}

func (r *PositionRepo) GetById(positionId string) (*models.Position, error) {
	var position *models.Position
	if err := r.db.First(position, "position_id = ?", positionId).Error; err != nil {
		return nil, utils.OnError(err, "an not get position by id")
	}
	return position, nil
}

func (r *PositionRepo) Create(position *models.Position) error {
	err := r.db.Create(&models.Position{}).Error
	return utils.OnError(err, "can not create position")
}

func (r *PositionRepo) Update(position *models.Position) error {
	if position.PositionId == 0 {
		return fmt.Errorf("can not update position: invalid input")
	}

	err := r.db.Save(position).Error
	return utils.OnError(err, "can not update position")
}

func (r *PositionRepo) DeleteById(positionId string) error {
	err := r.db.Where("position_id = ?", positionId).Delete(&models.Position{}).Error
	return utils.OnError(err, "can not delete position")
}
