package model_repo

import (
	"fmt"

	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func (r *OrderRepo) GetById(orderId string) (*models.Order, error) {
	order := &models.Order{}
	if err := r.db.First(order, "order_id = ?", orderId).Error; err != nil {
		return nil, utils.OnError(err, "can not find order")
	}
	return order, nil
}

func (r *OrderRepo) Create(order *models.Order) error {
	err := r.db.Create(order).Error
	return utils.OnError(err, "can not create order")
}

func (r *OrderRepo) Update(order *models.Order) error {
	if order.AccountId == 0 {
		return fmt.Errorf("can not update order: invalid input")
	}

	err := r.db.Save(order).Error
	return utils.OnError(err, "can not update order")
}

func (r *OrderRepo) DeleteById(orderId string) error {
	err := r.db.Where("order_id = ?", orderId).Delete(&models.Order{}).Error
	return utils.OnError(err, "can not delete order")
}
