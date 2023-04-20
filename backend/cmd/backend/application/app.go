package application

import (
	"context"
	"github.com/crypto-sign/cmd/backend/configuration"
	"github.com/crypto-sign/internal/services"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Config *configuration.Config
	Logger *logrus.Logger
}

func New(cfg *configuration.Config, logger *logrus.Logger) *App {
	return &App{
		Config: cfg,
		Logger: logger,
	}
}

func (a *App) Run() {
	envStructure := a.constructEnv()

	httpServer := a.NewHTTPServer(envStructure)

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				a.Logger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			a.Logger.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		a.Logger.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

type env struct {
	keyService *services.KeyService
	docService *services.DocService
}

func (a *App) constructEnv() *env {
	return &env{
		keyService: services.NewKeyService(a.Logger),
		docService: services.NewDocService(a.Logger),
	}
}
