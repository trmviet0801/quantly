package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SaveToMysqlDB[T any](db *gorm.DB, data *T) {
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(data)
}
