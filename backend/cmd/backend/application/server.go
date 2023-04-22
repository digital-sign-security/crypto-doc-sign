package application

import (
	"fmt"
	"github.com/crypto-sign/internal/handlers"
	"github.com/crypto-sign/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func (a *App) NewHTTPServer(env *env) *http.Server {
	mux := chi.NewMux()
	mux.Use(middleware.SignatureChecker)

	mux.Mount("/swagger", httpSwagger.WrapHandler)
	mux.Route("/v1", func(r chi.Router) {
		a.addDocsHandler(env, r)
		a.addKeysHandler(env, r)
	})

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler: mux,
	}
}

func (a *App) addDocsHandler(env *env, rg chi.Router) {
	handler := handlers.NewDocsHandler(env.docService)

	rg.Route("/docs", func(r chi.Router) {
		r.Post("/", handler.Create) // one handler for two modes
		r.Get("/available", handler.GetAvailable)
		r.Get("/{doc_id}", handler.Get)
	})
}

func (a *App) addKeysHandler(env *env, rg chi.Router) {
	handler := handlers.NewKeysHandler(env.keyService)

	rg.Route("/keys", func(r chi.Router) {
		r.Post("/public", handler.Post)
		r.Get("/public", handler.Get)
		r.Get("/public/{user_id}", handler.GetAnotherUserPublicKey)
		r.Get("/public/server", handler.GetServerPublicKey)

		r.Get("/", handler.GetGeneratedKeys)
	})
}
