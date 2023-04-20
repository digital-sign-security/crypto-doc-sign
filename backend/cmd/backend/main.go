package main

import (
	"github.com/crypto-sign/cmd/backend/application"
	"github.com/crypto-sign/cmd/backend/configuration"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	//ctx := context.Background()

	cfg := &configuration.Config{}

	logger := configureLogger()

	app := application.New(cfg, logger)
	app.Run()
	logger.Info("application stopped")
	//if err = app.Run(); err != nil {
	//	logger.Fatal("application stopped with error")
	//} else {
	//	logger.Info("application stopped")
	//}

}

func configureLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.Out = os.Stdout
	log.SetOutput(logger.Writer())
	return logger
}
