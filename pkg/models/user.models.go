package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserTgID int64
	Name     string
	Age      int
	Email    string
	Verified bool
}
