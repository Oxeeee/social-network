package service

import (
	"log/slog"

	"github.com/Oxeeee/social-network/internal/config"
	"github.com/Oxeeee/social-network/internal/repo"
)

type Service interface {
}

type service struct {
	log  *slog.Logger
	cfg  *config.Config
	repo repo.Repo
}

func NewService(log *slog.Logger, cfg *config.Config, repo repo.Repo) Service {
	return service{
		log:  log,
		cfg:  cfg,
		repo: repo,
	}
}
