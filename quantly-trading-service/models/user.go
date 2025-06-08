package models

type User struct {
	UserId   int64  `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}
