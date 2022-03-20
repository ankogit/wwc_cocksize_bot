package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"local/wwc_cocksize_bot/configs"
	"local/wwc_cocksize_bot/pkg/models"
	"local/wwc_cocksize_bot/pkg/storage"
	"log"
)

type Bot struct {
	bot            *tgbotapi.BotAPI
	config         *config.IniConf
	version        string
	messages       config.Messages
	users          map[int64]models.UserData
	userRepository storage.UserRepository
}

func NewBot(bot *tgbotapi.BotAPI, con *config.IniConf, version string, messages config.Messages, userRepository storage.UserRepository) *Bot {
	return &Bot{bot: bot, config: con, version: version, messages: messages, users: make(map[int64]models.UserData), userRepository: userRepository}
}

func (b *Bot) Start() error {
	// Авторизация бота
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)
	//updates := bot.ListenForWebhook("/" + bot.Token)
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	// Бесконечно ждем апдейтов от сервера
	for update := range updates {
		switch {
		// Пришло обычное сообщение
		case update.Message != nil && update.Message.ViaBot == nil && !update.Message.IsCommand():
			b.sendWelcomeMessage(update.Message)
			break

		case update.Message != nil && update.Message.IsCommand():
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}

			break

		// Пришел inline запрос
		case update.InlineQuery != nil:
			b.handleInlineQuery(update.InlineQuery)
			break
		}
	}
}
