package model_repo

import (
	"database/sql"
	"errors"

	"github.com/trmviet0801/quantly/models"
)

type OrderRepo struct {
	db *sql.DB
}

func (orderRepo *OrderRepo) GetById(orderId int64) (*models.Order, error) {
	query := `
	SELECT 
		order_id,
		account_id,
		symbol,
		qty,
		notional,
		side,
		type,
		time_in_force,
		limit_price,
		stop_price,
		trail_price,
		trail_percent,
		extended_hours,
		client_order_id,
		order_class
	FROM orders
	WHERE order_id = ?`

	row := orderRepo.db.QueryRow(query, orderId)
	order := &models.Order{}
	err := row.Scan(
		&order.OrderId,
		&order.AccountId,
		&order.Symbol,
		&order.Quantity,
		&order.Notional,
		&order.Side,
		&order.Type,
		&order.TimeInForce,
		&order.LimitPrice,
		&order.StopPrice,
		&order.TrailPrice,
		&order.TrailPercent,
		&order.ExtendedHours,
		&order.ClientOrderId,
		&order.OrderClass,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found with ID: " + string(orderId))
		}
	}
	return order, nil
}

func (orderRepo *OrderRepo) Create(order *models.Order) error {
	query := `
	INSERT INTO orders (
		account_id,
		symbol,
		qty,
		notional,
		side,
		type,
		time_in_force,
		limit_price,
		stop_price,
		trail_price,
		trail_percent,
		extended_hours,
		client_order_id,
		order_class
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := orderRepo.db.Exec(query,
		order.OrderId,
		order.AccountId,
		order.Symbol,
		order.Quantity,
		order.Notional,
		order.Side,
		order.Type,
		order.TimeInForce,
		order.LimitPrice,
		order.StopPrice,
		order.TrailPrice,
		order.TrailPercent,
		order.ExtendedHours,
		order.ClientOrderId,
		order.OrderClass,
	)
	if err != nil {
		return errors.New("failed to create order: " + err.Error())
	}
	return nil
}

func (orderRepo *OrderRepo) Update(order *models.Order) error {
	query := `
	UPDATE orders SET
		account_id = ?,
		symbol = ?,
		qty = ?,
		notional = ?,
		side = ?,
		type = ?,
		time_in_force = ?,
		limit_price = ?,
		stop_price = ?,
		trail_price = ?,
		trail_percent = ?,
		extended_hours = ?,
		client_order_id = ?,
		order_class = ?
	WHERE order_id = ?`
	_, err := orderRepo.db.Exec(query,
		order.AccountId,
		order.Symbol,
		order.Quantity,
		order.Notional,
		order.Side,
		order.Type,
		order.TimeInForce,
		order.LimitPrice,
		order.StopPrice,
		order.TrailPrice,
		order.TrailPercent,
		order.ExtendedHours,
		order.ClientOrderId,
		order.OrderClass,
		order.OrderId)
	if err != nil {
		return errors.New("failed to update order: " + err.Error())
	}
	return nil
}

func (orderRepo *OrderRepo) DeleteById(orderId int64) error {
	query := `DELETE FROM orders WHERE order_id = ?`
	_, err := orderRepo.db.Exec(query, orderId)
	if err != nil {
		return errors.New("failed to delete order: " + err.Error())
	}
	return nil
}
