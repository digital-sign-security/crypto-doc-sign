package main

import (
	"context"
	"github.com/crypto-sign/cmd/backend/application"
	"github.com/crypto-sign/cmd/backend/configuration"
	"github.com/crypto-sign/docs"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	_ = context.Background()

	logger := configureLogger()

	cfg, err := configuration.GetConfig(logger)
	if err != nil {
		logger.Fatalf("cannot read config: %v", err)
		return
	}
	configureSwaggerInfo(cfg)

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

func configureSwaggerInfo(cfg *configuration.Config) {
	docs.SwaggerInfo.Title = cfg.Swagger.Title
	docs.SwaggerInfo.Version = cfg.Swagger.Version
	docs.SwaggerInfo.BasePath = cfg.Swagger.BasePath
	docs.SwaggerInfo.Host = cfg.Swagger.Host
	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes
}
