package service

import (
	"strconv"
	"strings"

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

func (s *BuyFromCryptobotService) checkTetherBalance() error {
	paymentService := PaymentService{}
	balance, err := paymentService.GetTetherBalance()
	if err != nil {
		return err
	}
	balanceInt, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		return &consts.CustomError{
			Message: consts.PARSE_STRING_ERROR.Message,
			Code:    consts.PARSE_STRING_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	maxBuyReqInt, err := strconv.ParseFloat(consts.BUY_TETHER_AMOUNT_LIST[2], 64)
	if err != nil {
		return &consts.CustomError{
			Message: consts.PARSE_STRING_ERROR.Message,
			Code:    consts.PARSE_STRING_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	if balanceInt <= maxBuyReqInt+2 {
		return consts.LOW_BALANCE_ERROR
	}
	return nil
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

	if err := s.checkTetherBalance(); err != nil {
		return []tgbotapi.Chattable{}, err
	}

	if update.CallbackQuery != nil && update.CallbackQuery.Data != consts.BUY_TETHER_DATA {
		amount := strings.ReplaceAll(strings.Trim(update.CallbackQuery.Data, consts.BUY_TETHER_DATA), "|", "")
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, consts.PAY_FOR_TETHER_MESSAGE+amount, tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					consts.PAY_BUTTON_KEYBOARD,
					consts.PAY_BUTTON_DATA,
				),
			),
		))
		res = append(res, editedMsg)
	} else {
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, consts.BUY_SELECT_AMOUNT_MESSAGE, tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					consts.BUY_TETHER_AMOUNT_LIST[0],
					consts.BUY_TETHER_DATA+"|"+consts.BUY_TETHER_AMOUNT_LIST[0],
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					consts.BUY_TETHER_AMOUNT_LIST[1],
					consts.BUY_TETHER_DATA+"|"+consts.BUY_TETHER_AMOUNT_LIST[1],
				),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					consts.BUY_TETHER_AMOUNT_LIST[2],
					consts.BUY_TETHER_DATA+"|"+consts.BUY_TETHER_AMOUNT_LIST[2],
				),
			),
		))
		res = append(res, editedMsg)
	}

	return res, nil
}
