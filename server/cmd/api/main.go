package main

import (
	"log/slog"
	"os"

	"github.com/Oxeeee/social-network/internal/app"
	"github.com/Oxeeee/social-network/internal/config"
	"github.com/Oxeeee/social-network/internal/db"
	"github.com/Oxeeee/social-network/internal/repo"
	"github.com/Oxeeee/social-network/internal/service"
	"github.com/Oxeeee/social-network/internal/transport/handlers"
	authmw "github.com/Oxeeee/social-network/internal/utils/authmiddleware"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("starting...")

	db := db.ConnectDatabase(*cfg).GetDB()
	repo := repo.NewRepo(db, log)
	service := service.NewService(log, cfg, repo)
	handlers := handlers.NewHandler(log, service)

	mw := authmw.NewAuthMiddleware(log, cfg)
	server := app.New(log, handlers, mw)
	server.Start()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
