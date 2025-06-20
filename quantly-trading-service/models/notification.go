package models

type Notification struct {
	NotificationId int64  `json:"notification_id" gorm:"primaryKey;autoIncrement"`
	UserId         string `json:"user_id" gorm:"foreignKey:UserId;references:UserId;not null;type:varchar(64)"`
	Content        string `json:"content" gorm:"not null"`
}
