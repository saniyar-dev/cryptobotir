package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/saniyar-dev/cryptobotir/pkg/consts"
)

type BuyFromCryptobotService struct{}

func (s *BuyFromCryptobotService) Buy(update tgbotapi.Update) ([]tgbotapi.Chattable, error) {
	var chatID int64
	// var username string
	if update.Message != nil {
		chatID = update.Message.Chat.ID
		// username = update.Message.Chat.UserName
	}
	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID
		// username = update.CallbackQuery.Message.Chat.UserName
	}
	var res []tgbotapi.Chattable

	res = append(res, tgbotapi.NewMessage(chatID, consts.BUY_POLICY_MESSAGE))
	res = append(res, tgbotapi.NewMessage(chatID, consts.BUY_INSTRUCTION_MESSAGE))

	msg := tgbotapi.NewMessage(chatID, consts.BUY_SELECT_CURRENCY_MESSAGE)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				consts.BUY_TETHER_KEYBOARD,
				consts.BUY_TETHER_DATA,
			),
		),
	)
	res = append(res, msg)
	return res, nil
}

func (s *BuyFromCryptobotService) BuyTether(update tgbotapi.Update) ([]tgbotapi.Chattable, error) {
	var chatID int64
	var msgID int
	// var username string
	if update.Message != nil {
		chatID = update.Message.Chat.ID
		msgID = update.Message.MessageID
		// username = update.Message.Chat.UserName
	}
	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID
		msgID = update.CallbackQuery.Message.MessageID
		// username = update.CallbackQuery.Message.Chat.UserName
	}
	var res []tgbotapi.Chattable
	editedMsg := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, "test", tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"test",
				"test",
			),
		),
	))
	res = append(res, editedMsg)
	return res, nil
}
