package main

import (
	"log"
	"net/http"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thegera4/go-social-api/internal/store"
)

// Contains the main configuration settings for our application.
type application struct {
	config  config
	store  store.Storage
}

// Configuration settings for our application.
type config struct {
	addr 	string
	db		dbConfig
}

// Configuration settings for our database.
type dbConfig struct {
	addr 					string
	maxOpenConns 	int
	maxIdleConns 	int
	maxIdleTime		string
}

// Method that returns a new server mux for our application.
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middleware for router.
	r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal through ctx.Done() that the request has timed out and 
	// further processing should be stopped.
  r.Use(middleware.Timeout(60 * time.Second))

	// Route is a helper method that allows you to group and add a prefix to all endpoints.
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

// Method to start the HTTP server with the specified configuration.
func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,		
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10, // should be less than WriteTimeout
		IdleTimeout: time.Minute * 2,
	}

	log.Printf("Server started on %s", app.config.addr)

	return srv.ListenAndServe()
}