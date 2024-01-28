package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func init() {
	registredCommands["help"] = (*Commander).Help
}

func (commander *Commander) Help(msg *tgbotapi.MessageConfig, inputMsg *tgbotapi.Message) error {
	msg.Text =
		"/help - help" + "\n" +
			"/list - list peoducts" + "\n" +
			"/get - get product by index"
	return nil
}
