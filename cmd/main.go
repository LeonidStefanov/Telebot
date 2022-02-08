package main

import (
	"home/leonid/Git/Pract/telegram_bot/pkg/database"
	"home/leonid/Git/Pract/telegram_bot/pkg/geolocation"
	"home/leonid/Git/Pract/telegram_bot/pkg/service"
	"home/leonid/Git/Pract/telegram_bot/pkg/telebot"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	botKey    = "2076166423:AAGSEzlCugL_WvKtNwCXYl5IZS8V2_4fwrg"
	IPinfoKey = "c3e8941c306ed1"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	geo := geolocation.NewIPinfo(time.Second * 15)

	database, err := database.NewBD("leonid:0000@/dbecho")
	if err != nil {
		log.Println(err)
		return
	}
	defer database.Close()

	svc := service.NewService(database)

	telegramBot := telebot.NewBot(bot, geo, svc)

	err = telegramBot.ConnectBot()
	if err != nil {
		log.Fatal(err)
	}

}
