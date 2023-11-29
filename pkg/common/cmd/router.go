package main

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func createRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return r
}
