package main

import (
	"github.com/asdine/storm/v3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"local/wwc_cocksize_bot/configs"
	"local/wwc_cocksize_bot/pkg/storage/stormDB"
	"local/wwc_cocksize_bot/pkg/telegram"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	chatRepository := stormDB.NewChatRepository(db)

	bot, err := tgbotapi.NewBotAPI(telegramkey)
	if err != nil {
		log.Panic("Wrong key:", telegramkey, err)
	}

	bot.Debug = cfg.Debug

	timeZone, _ := time.LoadLocation("Europe/Moscow")
	scheduler := cron.New(cron.WithLocation(timeZone))
	defer scheduler.Stop()

	go func() {
		telegramBot := telegram.NewBot(
			bot,
			config,
			cfg.Version,
			cfg.Messages,
			userRepository,
			chatRepository,
		)

		telegramBot.CronInit(scheduler)
		go func() {
			telegramBot.CronStart()
		}()

		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	//serverInstance := new(server.Server)
	//handlers := new(handler.Handler)
	//
	//if err := serverInstance.RunHttp(cfg.Port, handlers.InitRoutes()); err != nil {
	//	log.Fatalf("error http server: %s", err.Error())
	//}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
