package model

import "time"

type User struct {
	Id        int64  `gorm:"column:id;primary_key;auto_increment"`
	Username  string `gorm:"column:username;unique"`
	Password  string `gorm:"column:password_hash"`
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
