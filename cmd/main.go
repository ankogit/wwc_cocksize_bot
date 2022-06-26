package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/asdine/storm/v3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"local/wwc_cocksize_bot/configs"
	"local/wwc_cocksize_bot/pkg/auth"
	"local/wwc_cocksize_bot/pkg/service"
	"local/wwc_cocksize_bot/pkg/storage"
	"local/wwc_cocksize_bot/pkg/storage/postgresDB"
	"local/wwc_cocksize_bot/pkg/storage/stormDB"
	"local/wwc_cocksize_bot/pkg/telegram"
	"local/wwc_cocksize_bot/pkg/transport/grpc"
	grpc_handler "local/wwc_cocksize_bot/pkg/transport/grpc/handler"
	"local/wwc_cocksize_bot/pkg/transport/rest"
	"local/wwc_cocksize_bot/pkg/transport/rest/handler"
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

	dbPostgres, err := postgresDB.NewPostgresDB(postgresDB.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("cant to connect postgress: %s", err))
		return
	}
	defer dbPostgres.Close()

	userRepository := stormDB.NewUserRepository(db)
	chatRepository := stormDB.NewChatRepository(db)

	repositories := storage.NewRepositories(db, dbPostgres)

	tokenManager, err := auth.NewManager(cfg.AuthSecret)
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

	serverInstance := new(rest.Server)
	handlers := handler.NewHandler(services)

	go func() {
		if err := serverInstance.RunHttp(cfg.Port, handlers.InitRoutes()); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error http server: %s", err.Error())
		}
	}()

	grpcHandlers := grpc_handler.NewGrpcHandler(services)
	grpcHandlers.InitHandlers()

	logger := logrus.New()
	grpcInstance := grpc.NewServer(grpc.Deps{
		Logger:       logger,
		StatsHandler: &grpcHandlers.StatsHandler,
	})
	go func() {
		log.Println("Starting gRPC server")
		if err := grpcInstance.ListenAndServe(1337); err != nil {
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
