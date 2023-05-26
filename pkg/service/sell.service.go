package service

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/saniyar-dev/cryptobotir/pkg/consts"
	"github.com/saniyar-dev/cryptobotir/pkg/models"
)

type SellFromCryptobotService struct{}

func (s *SellFromCryptobotService) Sell(update tgbotapi.Update) ([]tgbotapi.Chattable, error) {
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

	res = append(res, tgbotapi.NewMessage(chatID, consts.SELL_POLICY_MESSAGE))
	res = append(res, tgbotapi.NewMessage(chatID, consts.SELL_INSTRUCTION_MESSAGE))

	msg := tgbotapi.NewMessage(chatID, consts.SELL_SELECT_CURRENCY_MESSAGE)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				consts.SELL_TETHER_KEYBOARD,
				consts.SELL_TETHER_DATA,
			),
		),
	)
	res = append(res, msg)

	return res, nil
}

func (s *SellFromCryptobotService) SellTether(update tgbotapi.Update) ([]tgbotapi.Chattable, error) {
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

	if update.CallbackQuery != nil && update.CallbackQuery.Data != consts.SELL_TETHER_DATA {
		paymentService := PaymentService{}

		amount := strings.ReplaceAll(strings.Trim(update.CallbackQuery.Data, consts.SELL_TETHER_DATA), "|", "")
		invoice, err := paymentService.ReceiveTether(models.User{UserTgID: update.CallbackQuery.From.ID}, amount)
		if err != nil {
			return []tgbotapi.Chattable{}, err
		}

		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, consts.PAY_FOR_TETHER_MESSAGE+amount, tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL(
					consts.PAY_BUTTON_KEYBOARD,
					invoice.PayUrl,
				),
			),
		))
		res = append(res, editedMsg)
	} else {
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, consts.SELL_SELECT_AMOUNT_MESSAGE, tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					consts.SELL_TETHER_AMOUNT_LIST[0],
					consts.SELL_TETHER_DATA+"|"+consts.SELL_TETHER_AMOUNT_LIST[0],
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					consts.SELL_TETHER_AMOUNT_LIST[1],
					consts.SELL_TETHER_DATA+"|"+consts.SELL_TETHER_AMOUNT_LIST[1],
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					consts.SELL_TETHER_AMOUNT_LIST[2],
					consts.SELL_TETHER_DATA+"|"+consts.SELL_TETHER_AMOUNT_LIST[2],
				),
			),
		))
		res = append(res, editedMsg)
	}
	return res, nil
}
