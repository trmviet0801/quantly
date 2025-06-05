package model_repo

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/trmviet0801/quantly/models"
)

type StopLossRepo struct {
	db *sql.DB
}

func (stopLossRepo *StopLossRepo) GetById(stopLossId int64) (*models.StopLoss, error) {
	query := `
	SELECT 
		stop_loss_id,
		order_id,
		stop_price,
		limit_price
		FROM stop_losses WHERE stop_loss_id = ?
	`
	row := stopLossRepo.db.QueryRow(query, stopLossId)
	var stopLoss = &models.StopLoss{}
	err := row.Scan(
		&stopLoss.StopLostId,
		&stopLoss.OrderId,
		&stopLoss.StopPrice,
		&stopLoss.LimitPrice,
	)
	if err != nil {
		return nil, errors.New("stop loss not found with ID: " + strconv.FormatInt(stopLossId, 10))
	}
	return stopLoss, nil
}

func (stopLossRepo *StopLossRepo) Create(stopLoss *models.StopLoss) error {
	query := `
	INSERT INTO stop_losses (order_id, stop_price, limit_price)
	VALUES (?, ?, ?)
	`
	_, err := stopLossRepo.db.Exec(query, stopLoss.OrderId, stopLoss.StopPrice, stopLoss.LimitPrice)
	if err != nil {
		return errors.New("failed to create stop loss: " + err.Error())
	}
	return nil
}

func (stopLossRepo *StopLossRepo) Update(stopLoss *models.StopLoss) error {
	query := `
	UPDATE stop_losses
	SET order_id = ?, stop_price = ?, limit_price = ?
	WHERE stop_loss_id = ?
	`
	_, err := stopLossRepo.db.Exec(query, stopLoss.OrderId, stopLoss.StopPrice, stopLoss.LimitPrice, stopLoss.StopLostId)
	if err != nil {
		return errors.New("failed to update stop loss: " + err.Error())
	}
	return nil
}

func (stopLossRepo *StopLossRepo) Delete(stopLossId int64) error {
	query := `
	DELETE FROM stop_losses
	WHERE stop_loss_id = ?
	`
	_, err := stopLossRepo.db.Exec(query, stopLossId)
	if err != nil {
		return errors.New("failed to delete stop loss: " + err.Error())
	}
	return nil
}
