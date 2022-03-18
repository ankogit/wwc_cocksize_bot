package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"local/wwc_cocksize_bot/pkg/models"
	"log"
	"time"
)

func (b *Bot) handleInlineQuery(query *tgbotapi.InlineQuery) {
	var user models.UserData

	if val, _ := b.userRepository.Get(query.From.ID); val.ID != 0 {
		user = val
		fmt.Println("USER")
		fmt.Println(user)
		if user.Time.Add(time.Hour * 8).Before(time.Now()) {
			user.CockSize = getNewCockSizeV2(query.From.ID)
			user.Time = time.Now()
			err := b.userRepository.Save(user)
			if err != nil {
				log.Panic(err)
			}
		}
	} else {
		user = models.UserData{ID: query.From.ID, CockSize: getNewCockSizeV2(query.From.ID), Time: time.Now()}
		err := b.userRepository.Save(user)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("NEW")
		fmt.Println(user)
	}

	var resources []interface{}

	cockSizeMessage := getCockSizeMessage(user.CockSize)
	resources = append(resources,
		tgbotapi.InlineQueryResultArticle{
			Type:  "article",
			ID:    query.ID,
			Title: "🍆 Узнать свой размер",
			InputMessageContent: tgbotapi.InputTextMessageContent{
				Text: cockSizeMessage},
			Description: "Размер вашего штуцера сегодня"})

	// Отправляем меню пользователю
	if _, err := b.bot.Request(tgbotapi.InlineConfig{
		InlineQueryID: query.ID,
		CacheTime:     0,
		IsPersonal:    true,
		Results:       resources}); err != nil {
		log.Println(err)
	}
}
