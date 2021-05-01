package cmd

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testMascotGaming/internal/auth"
	"testMascotGaming/internal/client"
	"testMascotGaming/internal/handler"
	"testMascotGaming/internal/repository"
	"testMascotGaming/internal/repository/postgres"
	"testMascotGaming/internal/service"
	"testMascotGaming/pkg/logger"
	"testMascotGaming/pkg/server"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading environment: %s", err)
	}

	err = ParseConfig()
	if err != nil {
		log.Fatalf("error parsing config: %s", err.Error())
	}

	zapLogger, err := logger.NewLogger(viper.GetString("logger.level"))
	if err != nil {
		log.Fatalf("error initializing logger: %s", err)
	}

	dbConfig := postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db := postgres.ConnectionToPostgres(dbConfig)
	manager := auth.NewManager("qwqwqw", "qwerty")
	clientObj := client.NewClient("https://api.c27staging.netgamecore.com/v1/")

	repositoryObj := repository.NewRepository(db, zapLogger)
	serviceObj := service.NewService(repositoryObj, zapLogger, manager, clientObj)
	handlerObj := handler.NewHandler(serviceObj, zapLogger, manager, clientObj)

	s := new(server.Server)
	addr := fmt.Sprintf("%s:%s", "", viper.GetString("server.port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = s.Listen(addr, handlerObj.GetRouter()); err != nil {
			log.Fatalf("error listenning server: %s", err)
		}
		log.Print("Run listen server...")
	}()

	<-quit

	err = s.Shutdown(context.Background())
	if err != nil {
		log.Fatalf("error shutting server: %s", err)
	}
}

func ParseConfig() error {
	viper.SetConfigFile("./internal/config/config.yaml")
	return viper.ReadInConfig()
}
