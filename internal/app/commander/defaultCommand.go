package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (commander *Commander) DefaultCommand(msg *tgbotapi.MessageConfig, inputMsg *tgbotapi.Message) {
	msg.Text = "You say: " + inputMsg.Text + "\nI don't know that command"
}
