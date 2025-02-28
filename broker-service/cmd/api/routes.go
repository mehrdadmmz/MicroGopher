package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allwoed to connect, what methods are allowed, what headers are allowed, etc
	// this is a middleware that will be used for all routes
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},                                   // allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                 // allow all methods
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // allow all headers
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true, // allow cookies
		MaxAge:           300,  // 300 seconds
	}))

	mux.Use(middleware.Heartbeat("/ping")) // add a heartbeat route to check if the server is up

	mux.Post("/", app.Broker) // add a route to the broker handler

	return mux
}
