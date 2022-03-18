package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"local/wwc_cocksize_bot/config"
	"local/wwc_cocksize_bot/pkg/models"
	"local/wwc_cocksize_bot/pkg/repositories"
	"log"
)

type Bot struct {
	bot            *tgbotapi.BotAPI
	config         *config.IniConf
	version        string
	users          map[int64]models.UserData
	userRepository repositories.UserRepository
}

func NewBot(bot *tgbotapi.BotAPI, con *config.IniConf, version string, userRepository repositories.UserRepository) *Bot {
	return &Bot{bot: bot, config: con, version: version, users: make(map[int64]models.UserData), userRepository: userRepository}
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
		case update.Message != nil:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi!\n\nI'm inline bot for share your cocksize, type @"+b.bot.Self.UserName+" in message field. \nv. "+b.version+"")
			b.bot.Send(msg)
			break
		// Пришел inline запрос
		case update.InlineQuery != nil:
			b.handleInlineQuery(update.InlineQuery)
			break
		}
	}
}
