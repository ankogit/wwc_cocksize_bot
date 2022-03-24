package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"local/wwc_cocksize_bot/pkg/models"
	"log"
	"regexp"
	"sort"
	"strings"
	"time"
)

func (b *Bot) handleInlineQuery(query *tgbotapi.InlineQuery) {
	var user models.UserData

	if val, _ := b.userRepository.Get(query.From.ID); val.ID != 0 {
		user = val
		if user.Time.Add(time.Hour * 8).Before(time.Now()) {
			user.CockSize = getNewCockSizeV2(query.From.ID)
			user.Time = time.Now()
			user.Username = query.From.UserName
			user.FirstName = query.From.FirstName
			user.LastName = query.From.LastName
			err := b.userRepository.Save(user)
			if err != nil {
				log.Panic(err)
			}
		}
	} else {
		user = models.UserData{ID: query.From.ID, Username: query.From.UserName, FirstName: query.From.FirstName, LastName: query.From.LastName, CockSize: getNewCockSizeV2(query.From.ID), Time: time.Now()}
		err := b.userRepository.Save(user)
		if err != nil {
			log.Panic(err)
		}
	}

	var resources []interface{}

	cockSizeMessage := getCockSizeMessage(user.CockSize)
	resources = append(resources,
		tgbotapi.InlineQueryResultArticle{
			Type:  "article",
			ID:    query.ID,
			Title: b.messages.InlineContentTitle,
			InputMessageContent: tgbotapi.InputTextMessageContent{
				Text: cockSizeMessage},
			Description: b.messages.InlineContentDescription})

	// Отправляем меню пользователю
	if _, err := b.bot.Request(tgbotapi.InlineConfig{
		InlineQueryID: query.ID,
		CacheTime:     0,
		IsPersonal:    true,
		Results:       resources}); err != nil {
		log.Println(err)
	}
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		b.SendWelcomeMessage(message.Chat.ID)
		return nil
	case "stats":
		if err := b.handleCommandStats(message.Chat.ID); err != nil {
			return err
		}
		return nil
	case "notifications__enable":
		if err := b.handleNotificationEnable(message); err != nil {
			return err
		}
		return nil
	case "notifications__disable":
		if err := b.handleNotificationDisable(message); err != nil {
			return err
		}
		return nil
	default:
		return errUnknownCommand
	}
}

func (b *Bot) SendWelcomeMessage(chatId int64) {
	msg := tgbotapi.NewMessage(chatId, "Hi!\n\nI'm inline bot for share your cocksize, type @"+b.bot.Self.UserName+" in message field. \nv. "+b.version+"")
	b.bot.Send(msg)
}
func (b *Bot) SendStatsMessage(chatId int64) error {
	var usersInfo []string
	var textMessage string
	if users, err := b.userRepository.All(); users != nil {
		if err != nil {
			return err
		}
		sort.SliceStable(users, func(i, j int) bool {
			return users[i].CockSize > users[j].CockSize
		})

		for i, user := range users {
			username := user.Username
			if len(username) == 0 {
				username = "Anonymous"
			}
			msg := fmt.Sprintf("%v [%v](tg://user?id=%v) : *%v cm*", emojiBySize(user.CockSize), user.Username, user.ID, user.CockSize)
			if i == 0 {
				msg += " 👑"
			}
			usersInfo = append(usersInfo, msg)
		}
	}

	if len(usersInfo) > 0 {
		textMessage = fmt.Sprintf("Stats for %v \n", time.Now().Format("02.01.2006")) + strings.Join(usersInfo, "\n")
	} else {
		textMessage = b.messages.NoStats
	}
	msg := tgbotapi.NewMessage(chatId, textMessage)
	msg.ParseMode = "MARKDOWN"
	msg.DisableNotification = true
	b.bot.Send(msg)
	return nil
}

func (b *Bot) handleCommandStats(chatId int64) error {
	return b.SendStatsMessage(chatId)
}
func (b *Bot) handleNotificationEnable(message *tgbotapi.Message) error {
	if message.Chat == nil || message.Chat.Title == "" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Notifications are allowed to conduct group chats")
		b.bot.Send(msg)
		return nil
	}
	lenCommand := len("notifications__enable")
	cronParam := message.Text[lenCommand+2:]

	var validID = regexp.MustCompile(`(@(annually|yearly|monthly|weekly|daily|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7})`)
	if !validID.MatchString(cronParam) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Cron string is not correct :(")
		b.bot.Send(msg)
		return nil
	}

	var chat = models.Chat{
		ID:               message.Chat.ID,
		Title:            message.Chat.Title,
		NotificationCron: cronParam,
		EntryId:          0,
	}
	if err := b.chatRepository.Save(chat); err != nil {
		return err
	}

	entryId, err := b.cronService.SetJob(&chat, cronParam)
	if err != nil {
		return err
	}
	chat.EntryId = entryId

	if err := b.chatRepository.Save(chat); err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "Notifications are installed for this chat 🔔")
	b.bot.Send(msg)
	return nil
}

func (b *Bot) handleNotificationDisable(message *tgbotapi.Message) error {
	if message.Chat == nil || message.Chat.Title == "" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Notifications are allowed to conduct group chats")
		b.bot.Send(msg)
		return nil
	}

	chat, _ := b.chatRepository.Get(message.Chat.ID)

	if (chat == (models.Chat{})) || (chat.NotificationCron == "" && chat.EntryId == 0) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Notifications are not already setting for this chat")
		b.bot.Send(msg)
		return nil
	}

	b.cronService.RemoveJob(&chat)
	chat.NotificationCron = ""
	chat.EntryId = 0

	if err := b.chatRepository.Save(chat); err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "Notifications are disabled for this chat 🔕")
	b.bot.Send(msg)
	return nil
}
