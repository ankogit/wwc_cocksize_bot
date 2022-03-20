package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"local/wwc_cocksize_bot/pkg/models"
	"log"
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
	case "stats":
		if err := b.handleCommandStats(message); err != nil {
			return err
		}
		return nil
	default:
		return errUnknownCommand
	}
}

func (b *Bot) sendWelcomeMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hi!\n\nI'm inline bot for share your cocksize, type @"+b.bot.Self.UserName+" in message field. \nv. "+b.version+"")
	b.bot.Send(msg)
}

func (b *Bot) handleCommandStats(message *tgbotapi.Message) error {
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
	msg := tgbotapi.NewMessage(message.Chat.ID, textMessage)
	msg.ParseMode = "MARKDOWN"
	b.bot.Send(msg)

	return nil
}
