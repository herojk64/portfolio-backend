// @title          Portfolio API
// @version        1.0
// @description    REST API for portfolio backend — projects, skills, and settings.
// @host           localhost:8000
// @BasePath       /api/v1
// @schemes        http

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/herojk64/portfolio-backend/internal/app"
	"github.com/herojk64/portfolio-backend/internal/config"
	"github.com/herojk64/portfolio-backend/internal/database"
	"github.com/herojk64/portfolio-backend/internal/database/sqlc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	queries := sqlc.New(db)

	httpEngine := app.New(cfg)
	app.Handle(httpEngine, queries)

	s := &http.Server{
		Addr:           ":" + cfg.App.Port,
		Handler:        httpEngine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
