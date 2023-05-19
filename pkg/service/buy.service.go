package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BuyFromCryptobotService struct{}

func (s *BuyFromCryptobotService) Buy(update tgbotapi.Update) ([]tgbotapi.MessageConfig, error) {
	var res []tgbotapi.MessageConfig
	return res, nil
}

func (s *BuyFromCryptobotService) buyTether(update tgbotapi.Update) ([]tgbotapi.MessageConfig, error) {
	var res []tgbotapi.MessageConfig
	return res, nil
}
