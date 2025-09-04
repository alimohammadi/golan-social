package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alimohammadi/golan-social.git/internal/store"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
	store  store.Storage
	db     dbConfing
}

type dbConfing struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	addr string
	db   dbConfing
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)

	r.Route(
		"/v1", func(r chi.Router) {
			r.Get("/health", app.healthCheckHandler)
		},
	)

	// posts
	// users
	// auth
	return r
}

func (app *application) run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Print("server is started at $s", app.config.addr)

	return srv.ListenAndServe()
}
