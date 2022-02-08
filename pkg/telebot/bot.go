package telebot

import (
	"home/leonid/Git/Pract/telegram_bot/pkg/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	geo geo

	svc service.Service
}

type geo interface {
	GetGeo(ip string) string
}

func NewBot(bot *tgbotapi.BotAPI, g geo, svc service.Service) *Bot {
	return &Bot{
		bot: bot,
		geo: g,
		svc: svc,
	}
}

func (b *Bot) ConnectBot() error {
	// log.Printf("%s  подключен ", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue

		}

		b.handleIPinfo(update.Message)
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)

}
