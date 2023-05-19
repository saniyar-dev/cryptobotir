package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserTgID string
	Name     string
	Age      int
	Email    string
	Verified bool
}
