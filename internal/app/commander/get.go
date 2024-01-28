package commander

import (
	"errors"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	registredCommands["get"] = (*Commander).Get
}

func (commander *Commander) Get(msg *tgbotapi.MessageConfig, inputMsg *tgbotapi.Message) error {
	args := inputMsg.CommandArguments()

	num, errParseNum := strconv.ParseInt(args, 10, 64)
	if errParseNum != nil {
		errText := "Error. Can`t take \"" + args + "\" elememt. Incorrect param of command \"get\""
		return errors.New(errText)
	}

	foundElem, errGet := commander.productService.Get(num)
	if errGet != nil {
		return errGet
	}

	msg.Text = "The " + args + " element:\n" + foundElem.Title
	return nil
}
