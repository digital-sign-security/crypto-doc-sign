package main

import (
	"context"
	"github.com/crypto-sign/cmd/backend/application"
	"github.com/crypto-sign/cmd/backend/configuration"
	"github.com/crypto-sign/docs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	_ = context.Background()

	cfg := &configuration.Config{}

	if err := configuration.InitConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}
	configureSwaggerInfo()

	logger := configureLogger()

	app := application.New(cfg, logger)
	if err := app.Run(); err != nil {
		logger.Fatal("application stopped with error")
	} else {
		logger.Info("application stopped")
	}

}

func configureLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.Out = os.Stdout
	log.SetOutput(logger.Writer())
	return logger
}

func configureSwaggerInfo() {
	docs.SwaggerInfo.Title = viper.GetString("swagger.title")
	docs.SwaggerInfo.Version = viper.GetString("swagger.version")
	docs.SwaggerInfo.BasePath = viper.GetString("swagger.base-path")
	docs.SwaggerInfo.Host = viper.GetString("swagger.host")
	docs.SwaggerInfo.Schemes = viper.GetStringSlice("swagger.schemes")
}
