package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"local/wwc_cocksize_bot/configs"
	"local/wwc_cocksize_bot/pkg/models"
	"local/wwc_cocksize_bot/pkg/service"
	"local/wwc_cocksize_bot/pkg/storage"
	"local/wwc_cocksize_bot/pkg/telegram/jobs"
	"log"
)

type Bot struct {
	bot            *tgbotapi.BotAPI
	config         *config.IniConf
	version        string
	messages       config.Messages
	users          map[int64]models.UserData
	userRepository storage.UserRepository
	chatRepository storage.ChatRepository
	cronService    *service.CronService
}

func NewBot(bot *tgbotapi.BotAPI, con *config.IniConf, version string, messages config.Messages, userRepository storage.UserRepository, charRepository storage.ChatRepository) *Bot {
	return &Bot{bot: bot, config: con, version: version, messages: messages, users: make(map[int64]models.UserData), userRepository: userRepository, chatRepository: charRepository}
}

func (b *Bot) CronInit(scheduler *cron.Cron) {
	telegramNotifications := jobs.NewTelegramNotifications(b)
	//telegramNotifications := jobs.TelegramNotifications{BotApi}
	b.cronService = service.NewCronService(scheduler, b.chatRepository, telegramNotifications)
	b.cronService.Init()
}
func (b *Bot) CronStart() {
	b.cronService.Start()
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
		case update.Message != nil && update.Message.ViaBot == nil && !update.Message.IsCommand() && update.Message.ReplyToMessage == nil && update.Message.Chat.Type == "private":
			b.SendWelcomeMessage(update.Message.Chat.ID)
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
