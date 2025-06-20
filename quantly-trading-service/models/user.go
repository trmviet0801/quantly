package models

type User struct {
	UserId   string `json:"user_id" gorm:"primaryKey;autoIncrement:type:varchar(64)"`
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}
