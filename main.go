package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"time"
)

const version = "0.0.7"

func main() {
	var users = make(map[int64]UserData)
	var user UserData

	config := new(IniConf)
	config.CheckAndLoadConf("config" + string(os.PathSeparator) + "config.ini")
	telegramkey := config.GetStringKey("", "telegramkey")

	bot, err := tgbotapi.NewBotAPI(telegramkey)
	if err != nil {
		log.Panic("Wrong key:", telegramkey, err)
	}

	bot.Debug = config.GetBoolKey("", "debug")

	// Авторизация бота
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	//updates := bot.ListenForWebhook("/" + bot.Token)

	// Бесконечно ждем апдейтов от сервера
	for update := range updates {
		switch {
		// Пришло обычное сообщение
		case update.Message != nil:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi!\n\nI'm inline bot for share your cocksize, type @"+bot.Self.UserName+" in message field. \nv. "+version+"")
			bot.Send(msg)
			break
		// Пришел inline запрос
		case update.InlineQuery != nil:

			if val, ok := users[update.InlineQuery.From.ID]; ok {
				user = val

				if user.time.Add(time.Hour * 24).Before(time.Now()) {
					user.cockSize = getNewCockSize()
					user.time = time.Now()
				}
			} else {
				user = UserData{getNewCockSize(), time.Now()}
				users[update.InlineQuery.From.ID] = user
			}

			var resources []interface{}

			cockSizeMessage := getCockSizeMessage(user.cockSize)
			resources = append(resources,
				tgbotapi.InlineQueryResultArticle{
					Type:  "article",
					ID:    update.InlineQuery.ID,
					Title: "🍆 Узнать свой размер",
					InputMessageContent: tgbotapi.InputTextMessageContent{
						Text: cockSizeMessage},
					Description: "Поделится размер штуцера сегодня"})

			// Отправляем меню пользователю
			if _, err := bot.Request(tgbotapi.InlineConfig{
				InlineQueryID: update.InlineQuery.ID,
				CacheTime:     0,
				Results:       resources}); err != nil {
				log.Println(err)
			}
			break
		}
	}
}
