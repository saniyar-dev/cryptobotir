package models

import (
	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	UserID    int
	Amount    int
	Currency  cryptobot.Currency
	IRLAmount int
}
