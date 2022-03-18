package main

import (
	"github.com/asdine/storm/v3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"local/wwc_cocksize_bot/config"
	"local/wwc_cocksize_bot/pkg/repositories/stormDB"
	"local/wwc_cocksize_bot/pkg/telegram"
	"log"
	"os"
)

const version = "0.3.1"

func main() {
	config := new(config.IniConf)
	config.CheckAndLoadConf("config" + string(os.PathSeparator) + "config.ini")
	telegramkey := config.GetStringKey("", "telegramkey")

	db, err := storm.Open(config.GetStringKey("", "dbName"))
	defer db.Close()

	userRepository := stormDB.NewUserRepository(db)

	bot, err := tgbotapi.NewBotAPI(telegramkey)
	if err != nil {
		log.Panic("Wrong key:", telegramkey, err)
	}

	bot.Debug = config.GetBoolKey("", "debug")

	telegramBot := telegram.NewBot(bot, config, version, userRepository)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
