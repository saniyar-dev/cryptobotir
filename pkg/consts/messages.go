package consts

var (
	START_COMMAND = "start"
	HELP_COMMAND  = "help"
	BUY_COMMAND   = "buy"
	SELL_COMMAND  = "sell"

	START_MESSAGE                = "start"
	HELP_MESSAGE                 = "help"
	BUY_INSTRUCTION_MESSAGE      = "buy instruction"
	BUY_POLICY_MESSAGE           = "buy policy"
	BUY_SELECT_CURRENCY_MESSAGE  = "select currency"
	BUY_SELECT_AMOUNT_MESSAGE    = "select amount in USD"
	SELL_INSTRUCTION_MESSAGE     = "sell instruction"
	SELL_POLICY_MESSAGE          = "sell policy"
	SELL_SELECT_CURRENCY_MESSAGE = "select currency"
	SELL_SELECT_AMOUNT_MESSAGE   = "select amount in USD"

	BUY_TETHER_KEYBOARD     = "Tether"
	BUY_TETHER_DATA         = "Buy Tether"
	BUY_TETHER_AMOUNT_LIST  = []string{"2", "4", "6"}
	SELL_TETHER_KEYBOARD    = "Tether"
	SELL_TETHER_DATA        = "Sell Tether"
	SELL_TETHER_AMOUNT_LIST = []string{"10", "15", "25"}
	PAY_FOR_TETHER_MESSAGE  = "pay for tether! Amount: "

	PAY_BUTTON_KEYBOARD = "go to pay!"
	PAY_BUTTON_DATA     = "go to pay!"
)
