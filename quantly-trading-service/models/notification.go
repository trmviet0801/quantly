package models

type Notification struct {
	NotificationId int64  `json:"notification_id" gorm:"primaryKey;autoIncrement"`
	UserId         int64  `json:"user_id" gorm:"foreignKey:UserId;references:UserId;not null"`
	Content        string `json:"content" gorm:"not null"`
}
