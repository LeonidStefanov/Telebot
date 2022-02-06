package telebot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart  = "start"
	commandStatus = "status"
	commandHеlp   = "help"
	commandMe     = "me"
)

func (b *Bot) handleMassage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.bot.Send(msg)

}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды :(")

	switch message.Command() {
	case commandStart:
		msg.Text = "Привет!Я testIPinfo_bot"

		_, err := b.bot.Send(msg)

		return err
	case commandHеlp:
		msg.Text = "Вот такие команды я знаю: /status  /me"
		_, err := b.bot.Send(msg)

		return err

	case commandStatus:
		msg.Text = "Я в порядке) Слежу за тобой 👀"
		_, err := b.bot.Send(msg)

		return err
	case commandMe:
		name := message.From.FirstName
		nickname := message.From.UserName
		id := message.From.ID

		msg.Text = fmt.Sprintf("Твоё имя: %v Ник: %v Id:%d", name, nickname, id)

		_, err := b.bot.Send(msg)
		return err

	default:
		_, err := b.bot.Send(msg)

		return err

	}

}
