package commander

import (
	"github.com/fahretD/telegram_bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var registredCommands = map[string]func(commander *Commander, msg *tgbotapi.MessageConfig, inputMsg *tgbotapi.Message) error{}

type Commander struct {
	Bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{bot, productService}
}

func (commander *Commander) HandleUpdates(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if command, ok := registredCommands[update.Message.Command()]; ok {
		err := command(commander, &msg, update.Message)
		if err != nil {
			msg.Text = err.Error()
		}
	} else {
		commander.DefaultCommand(&msg, update.Message)
	}
	commander.Bot.Send(msg)
}
