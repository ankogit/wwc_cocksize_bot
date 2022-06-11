package main

import (
	"context"
	"errors"
	"github.com/asdine/storm/v3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"local/wwc_cocksize_bot/configs"
	"local/wwc_cocksize_bot/pkg/auth"
	"local/wwc_cocksize_bot/pkg/server"
	"local/wwc_cocksize_bot/pkg/server/handler"
	"local/wwc_cocksize_bot/pkg/service"
	"local/wwc_cocksize_bot/pkg/storage"
	"local/wwc_cocksize_bot/pkg/storage/stormDB"
	"local/wwc_cocksize_bot/pkg/telegram"
	"log"
	"net/http"
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

	repositories := storage.NewRepositories(db)

	tokenManager, err := auth.NewManager("secretKey")
	if err != nil {
		log.Fatal(err)
		return
	}
	services := service.NewServices(service.Deps{
		Repositories: repositories,
		TokenManager: tokenManager,
	})

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
			services,
		)

		telegramBot.CronInit(scheduler)
		go func() {
			telegramBot.CronStart()
		}()

		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	serverInstance := new(server.Server)
	handlers := handler.NewHandler(services)

	go func() {
		if err := serverInstance.RunHttp(cfg.Port, handlers.InitRoutes()); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	log.Println("Started.")
	<-quit

	if err := serverInstance.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occurated on shuting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error occurated on db close: %s", err.Error())
	}
	log.Println("Shutdown...")
}
