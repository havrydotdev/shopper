package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"shopper"
	"shopper/pkg/handler"
	"shopper/pkg/repo"
	"shopper/pkg/service"
	"syscall"
)

// @title ShopperGo
// @version 1.0
// @description Open-source API project for online shop

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.TextFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error occures while initializing config file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("errorResponse occured while initializing server, error: %s", err.Error())
	}

	postgresDB, err := repo.NewPostgresDB(repo.Config{
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		Password: os.Getenv("DB_PASSWORD"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Error occured while initializing server, error: %s", err.Error())
	}

	repos := repo.NewRepository(postgresDB)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(shopper.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while starting server: %s", err.Error())
		}
	}()

	logrus.Println("ShopperGo started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("ShopperGo is shutting down")

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := postgresDB.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
