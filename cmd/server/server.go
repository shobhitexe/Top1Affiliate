package main

import (
	"log"
	"net/http"
	"top1affiliate/internal/routes"
	"top1affiliate/pkg/di"

	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

type APIServer struct {
	config Config
	db     *pgxpool.Pool
}

type Config struct {
	Addr     string
	dbConfig DBConfig
}

type DBConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (s *APIServer) mount() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	container := di.NewContainer(s.db)

	r.Route("/api/v1", func(r chi.Router) {
		routes.RegisterDataRoutes(r, container.DataHandler)
		routes.RegisterUserRoutes(r, container.UserHandler)
	})

	return r

}

func (s *APIServer) run(mux http.Handler) error {
	srv := http.Server{
		Addr:         s.config.Addr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("API server started at port: %s", srv.Addr)

	return srv.ListenAndServe()
}
