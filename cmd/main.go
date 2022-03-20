package main

import (
	"github.com/asdine/storm/v3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"local/wwc_cocksize_bot/configs"
	"local/wwc_cocksize_bot/pkg/storage/stormDB"
	"local/wwc_cocksize_bot/pkg/telegram"
	"log"
	"os"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Panic(err)
	}

	config := new(config.IniConf)
	config.CheckAndLoadConf("configs" + string(os.PathSeparator) + "config.ini")
	telegramkey := config.GetStringKey("", "telegramkey")

	db, err := storm.Open(config.GetStringKey("", "dbName"))
	defer db.Close()

	userRepository := stormDB.NewUserRepository(db)

	bot, err := tgbotapi.NewBotAPI(telegramkey)
	if err != nil {
		log.Panic("Wrong key:", telegramkey, err)
	}

	bot.Debug = cfg.Debug

	telegramBot := telegram.NewBot(bot, config, cfg.Version, cfg.Messages, userRepository)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
