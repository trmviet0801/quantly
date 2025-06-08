package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func (r *OrderRepo) GetById(orderId string) (*models.Order, error) {
	var order *models.Order
	err := r.db.First(order, "order_id = ?", orderId).Error
	if err != nil {
		return nil, fmt.Errorf("can not find order: %w", err)
	}
	return order, nil
}

func (r *OrderRepo) Create(order *models.Order) error {
	err := r.db.Create(order).Error
	if err != nil {
		return fmt.Errorf("can not create order: %w", err)
	}
	return nil
}

func (r *OrderRepo) Update(order *models.Order) error {
	if order.AccountId == 0 {
		return fmt.Errorf("can not update order: invalid input")
	}
	err := r.db.Save(order).Error
	if err != nil {
		return fmt.Errorf("can not update order: %w", err)
	}
	return nil
}

func (r *OrderRepo) DeleteById(orderId string) error {
	err := r.db.Where("order_id = ?", orderId).Delete(&models.Order{}).Error
	if err != nil {
		return fmt.Errorf("can not delete order: %w", err)
	}
	return nil
}
