package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	errUnknownCommand = errors.New("unknown command")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case errUnknownCommand:
		messageText = b.messages.Errors.UnknownCommand
	default:
		messageText = err.Error()
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}
