package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"shopper"
	"shopper/pkg/handler"
	"shopper/pkg/repo"
	"shopper/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.TextFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error occures while initializing config file: %s", err.Error())
	}

	postgresDB, err := repo.NewPostgresDB(repo.Config{
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		Password: viper.GetString("db.password"),
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
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while starting server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
