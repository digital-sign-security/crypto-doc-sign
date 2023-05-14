package application

import (
	"context"
	"github.com/crypto-sign/cmd/backend/configuration"
	"github.com/crypto-sign/internal/clients"
	"github.com/crypto-sign/internal/generators"
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

func (a *App) Run() error {
	envStructure := a.constructEnv()

	httpServer := a.NewHTTPServer(envStructure)

	postgresClient, err := clients.NewClient(context.TODO(), a.Config.Storage, a.Logger)
	if err != nil {
		a.Logger.Fatal(err)
		return err
	}
	a.Logger.Infof("%v", postgresClient)

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

	a.Logger.Info("server was started successfully")
	// Run the server
	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		a.Logger.Fatal(err)
		return err
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
	return nil
}

type env struct {
	keyService  *services.KeyService
	docService  *services.DocService
	userService *services.UserService
}

func (a *App) constructEnv() *env {
	generator := generators.NewKeysGenerator(a.Config.Generator)

	return &env{
		keyService:  services.NewKeyService(a.Logger, generator),
		docService:  services.NewDocService(a.Logger),
		userService: services.NewUserService(a.Logger),
	}
}
