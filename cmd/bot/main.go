package main

import (
	"log"
	"os"
	"telegram_bot/internal/app/commander"
	"telegram_bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, errr := tgbotapi.NewBotAPI(token)
	if errr != nil {
		log.Panic(errr)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()

	commander := commander.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdates(&update)
	}
}
