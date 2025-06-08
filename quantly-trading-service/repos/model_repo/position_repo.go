package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type PositionRepo struct {
	db *gorm.DB
}

func (r *PositionRepo) GetById(positionId string) (*models.Position, error) {
	var position *models.Position
	err := r.db.First(position, "position_id = ?", positionId).Error
	if err != nil {
		return nil, fmt.Errorf("can not get position by id: %w", err)
	}
	return position, nil
}

func (r *PositionRepo) Create(position *models.Position) error {
	err := r.db.Create(&models.Position{}).Error
	if err != nil {
		return fmt.Errorf("can not create position: %w", err)
	}
	return nil
}

func (r *PositionRepo) Update(position *models.Position) error {
	if position.PositionId == 0 {
		return fmt.Errorf("can not update position: invalid input")
	}
	err := r.db.Save(position).Error
	if err != nil {
		return fmt.Errorf("can not update position: %w", err)
	}
	return nil
}

func (r *PositionRepo) DeleteById(positionId string) error {
	err := r.db.Where("position_id = ?", positionId).Delete(&models.Position{}).Error
	if err != nil {
		return fmt.Errorf("can not delete position: %w", err)
	}
	return nil
}
