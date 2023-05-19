package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/saniyar-dev/cryptobotir/pkg/consts"
	"github.com/saniyar-dev/cryptobotir/pkg/service"
)

type MessageHandler struct{}

func (h *MessageHandler) handleCommands(update tgbotapi.Update) ([]tgbotapi.MessageConfig, error) {
	var res []tgbotapi.MessageConfig
	var err error
	chatID := update.Message.Chat.ID
	switch update.Message.Command() {
	case consts.START_COMMAND:
		res = append(res, tgbotapi.NewMessage(chatID, consts.START_MESSAGE))
	case consts.HELP_COMMAND:
		res = append(res, tgbotapi.NewMessage(chatID, consts.HELP_MESSAGE))
	case consts.BUY_COMMAND:
		buyService := service.BuyFromCryptobotService{}
		res, err = buyService.Buy(update)
		if err != nil {
			return []tgbotapi.MessageConfig{}, err
		}
	default:
		return []tgbotapi.MessageConfig{}, consts.UPDATE_MESSAGE_ERROR
	}
	return res, nil
}

func (h *MessageHandler) handleCallbackQuery(
	update tgbotapi.Update,
) ([]tgbotapi.MessageConfig, error) {
	var res []tgbotapi.MessageConfig
	// var err error

	switch update.CallbackQuery.Data {
	default:
		return []tgbotapi.MessageConfig{}, &consts.CustomError{
			Message: consts.UPDATE_MESSAGE_ERROR.Message,
			Code:    consts.UPDATE_MESSAGE_ERROR.Code,
			Detail:  update.CallbackQuery.Data,
		}
	}
	return res, nil
}

func (h *MessageHandler) HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	var res []tgbotapi.MessageConfig
	var err error

	if update.Message != nil {
		if update.Message.IsCommand() {
			res, err = h.handleCommands(update)
			if err != nil {
				return err
			}
		} else {
			return consts.UPDATE_MESSAGE_ERROR
		}
	}

	if update.CallbackQuery != nil {
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		if _, err := bot.Request(callback); err != nil {
			return &consts.CustomError{
				Message: consts.BOT_HANDLE_CALLBACKQUERY_ERROR.Message,
				Code:    consts.BOT_HANDLE_CALLBACKQUERY_ERROR.Code,
				Detail:  err.Error(),
			}
		}

		res, err = h.handleCallbackQuery(update)
		if err != nil {
			return err
		}
	}

	for _, msg := range res {
		if _, err := bot.Send(msg); err != nil {
			return &consts.CustomError{
				Message: consts.BOT_SEND_ERROR.Message,
				Code:    consts.BOT_SEND_ERROR.Code,
				Detail:  err.Error(),
			}
		}
	}

	return nil
}
