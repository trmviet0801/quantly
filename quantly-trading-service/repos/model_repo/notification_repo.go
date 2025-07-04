package model_repo

import (
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/utils"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	DB *gorm.DB
}

func (r *NotificationRepo) GetById(notificationId int64) (*models.Notification, error) {
	notification := &models.Notification{}
	if err := r.DB.First(notification, "notification_id = ?", notificationId).Error; err != nil {
		return nil, utils.OnError(err, "can not get notification")
	}

	return notification, nil
}

func (r *NotificationRepo) Create(notification *models.Notification) error {
	err := r.DB.Create(notification).Error
	return utils.OnError(err, "can not create notification")
}

func (r *NotificationRepo) Update(notification *models.Notification) error {
	if notification.NotificationId == 0 {
		return gorm.ErrRecordNotFound
	}

	err := r.DB.Save(notification).Error
	return utils.OnError(err, "can not update notification")
}

func (r *NotificationRepo) DeleteById(notificationId int64) error {
	err := r.DB.Where("notification_id = ?", notificationId).Delete(&models.Notification{}).Error
	return utils.OnError(err, "can not delete notification")
}
