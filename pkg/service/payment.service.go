package service

import (
	"fmt"
	"log"
	"os"

	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
	"github.com/saniyar-dev/cryptobotir/pkg/consts"
	"github.com/saniyar-dev/cryptobotir/pkg/models"
)

type PaymentService struct{}

var paymentClient *cryptobot.Client

func (p *PaymentService) InitPaymentClient() error {
	testMode := os.Getenv("CRYPTO_BOT_TEST_MODE")
	apiToken := os.Getenv("CRYPTO_BOT_API_TOKEN")
	paymentClient = cryptobot.NewClient(cryptobot.Options{
		APIToken: apiToken,
		Testing:  testMode == "true",
	})

	appInfo, err := paymentClient.GetMe()
	if err != nil {
		return &consts.CustomError{
			Message: consts.CRYPTO_BOT_CRASH_ERROR.Message,
			Code:    consts.CRYPTO_BOT_CRASH_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	log.Printf(
		"Crypto App info: AppID - %v, Name - %s, PaymentProcessingBotUsername - %s \n",
		appInfo.AppID,
		appInfo.Name,
		appInfo.PaymentProcessingBotUsername,
	)

	return nil
}

func (s *PaymentService) ReceiveTether(user models.User, amount string) (*cryptobot.Invoice, error) {
	invoice, err := paymentClient.CreateInvoice(cryptobot.CreateInvoiceRequest{
		Asset:          cryptobot.USDT,
		Amount:         amount,
		Description:    "",
		HiddenMessage:  "",
		PaidBtnName:    "callback",
		PaidBtnUrl:     "https://t.me/ircryptobuybot",
		Payload:        "",
		AllowComments:  true,
		AllowAnonymous: false,
		ExpiresIn:      60 * 10,
	})
	if err != nil {
		return &cryptobot.Invoice{}, &consts.CustomError{
			Message: consts.CRYPTO_BOT_CREATE_INVOICE_ERROR.Message,
			Code:    consts.CRYPTO_BOT_CREATE_INVOICE_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	return invoice, nil
}

func (s *PaymentService) TransferTether(user models.User, amount string) error {
	transfer, err := paymentClient.Transfer(cryptobot.TransferRequest{
		UserID:                  user.UserTgID,
		Asset:                   cryptobot.USDT,
		Amount:                  amount,
		SpendID:                 "",
		Comment:                 "Debt",
		DisableSendNotification: false,
	})
	if err != nil {
		return &consts.CustomError{
			Message: consts.CRYPTO_BOT_TRANSFER_ERROR.Message,
			Code:    consts.CRYPTO_BOT_TRANSFER_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	fmt.Printf(
		"ID - %v, UserID - %s, Status - %s, Amount - %s, Asset - %s, Comment - %s, CompletedAt - %s \n",
		transfer.ID,
		transfer.UserID,
		transfer.Status,
		transfer.Amount,
		transfer.Asset,
		transfer.Comment,
		transfer.CompletedAt,
	)
	return nil
}
