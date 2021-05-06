package cmd

import (
	"context"
	"fmt"
	"github.com/EgorMizerov/testMascotGaming/internal/auth"
	"github.com/EgorMizerov/testMascotGaming/internal/client"
	"github.com/EgorMizerov/testMascotGaming/internal/handler"
	"github.com/EgorMizerov/testMascotGaming/internal/repository"
	"github.com/EgorMizerov/testMascotGaming/internal/repository/postgres"
	"github.com/EgorMizerov/testMascotGaming/internal/service"
	"github.com/EgorMizerov/testMascotGaming/pkg/logger"
	"github.com/EgorMizerov/testMascotGaming/pkg/server"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	viper.SetConfigFile("./config.yaml")
	return viper.ReadInConfig()
}
