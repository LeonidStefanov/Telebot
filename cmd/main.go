package main

import (
	"home/leonid/Git/Pract/telegram_bot/pkg/telebot"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const botKey = "2076166423:AAGSEzlCugL_WvKtNwCXYl5IZS8V2_4fwrg"

func main() {

	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telebot.NewBot(bot)

	err = telegramBot.ConnectBot()
	if err != nil {
		log.Fatal(err)
	}

}
