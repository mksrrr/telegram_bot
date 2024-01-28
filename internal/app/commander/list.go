package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func init() {
	registredCommands["list"] = (*Commander).List
}

func (commander *Commander) List(msg *tgbotapi.MessageConfig, inputMsg *tgbotapi.Message) error {
	msg.Text = "All products:\n"

	products := commander.productService.List()
	for i, p := range products {
		msg.Text += "- " + p.Title
		if i != len(products)-1 {
			msg.Text += "\n"
		}
	}
	return nil
}
