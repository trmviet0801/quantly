package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	db *gorm.DB
}

func (r *NotificationRepo) GetById(notificationId int64) (*models.Notification, error) {
	var notification *models.Notification
	err := r.db.First(notification, "notification_id = ?", notificationId).Error
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *NotificationRepo) Create(notification *models.Notification) error {
	err := r.db.Create(notification).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationRepo) Update(notification *models.Notification) error {
	if notification.NotificationId == 0 {
		return gorm.ErrRecordNotFound
	}
	err := r.db.Save(notification).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationRepo) DeleteById(notificationId int64) error {
	err := r.db.Where("notification_id = ?", notificationId).Delete(&models.Notification{}).Error
	if err != nil {
		return err
	}
	return nil
}
