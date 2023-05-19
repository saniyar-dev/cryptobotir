package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TgID   string
	ChatID string
}
