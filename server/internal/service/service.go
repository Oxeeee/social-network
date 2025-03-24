package service

import (
	"log/slog"

	"github.com/Oxeeee/social-network/internal/config"
	"github.com/Oxeeee/social-network/internal/models/domain"
	"github.com/Oxeeee/social-network/internal/models/requests"
	"github.com/Oxeeee/social-network/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(req requests.Register) error
}

type service struct {
	log  *slog.Logger
	cfg  *config.Config
	repo repo.Repo
}

func NewService(log *slog.Logger, cfg *config.Config, repo repo.Repo) Service {
	return &service{
		log:  log,
		cfg:  cfg,
		repo: repo,
	}
}

func (s *service) Register(req requests.Register) error {
	const op = "service.register"
	log := s.log.With(slog.String("op", op))
	var user domain.User = domain.User{
		Email:    req.Email,
		Username: req.Username,
		Name:     req.Name,
		Surname:  req.Surname,
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error while generating hash", "error", err)
		return err
	}

	user.PassHash = string(hashPass)

	err = s.repo.Register(user)
	if err != nil {
		return err
	}

	return nil
}
