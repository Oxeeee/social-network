package service

import (
	"log/slog"

	"github.com/Oxeeee/social-network/internal/config"
	"github.com/Oxeeee/social-network/internal/models/domain"
	cerrors "github.com/Oxeeee/social-network/internal/models/errors"
	"github.com/Oxeeee/social-network/internal/models/requests"
	"github.com/Oxeeee/social-network/internal/repo"
	"github.com/Oxeeee/social-network/internal/utils/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(req requests.Register) error
	Login(req requests.Login) (string, string, error)
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

func (s *service) Login(req requests.Login) (string, string, error) {
	const op = "service.login"
	log := s.log.With(slog.String("op", op))

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(req.Password)); err != nil {
		log.Debug("compare hash and password", "error", err)
		return "", "", cerrors.ErrInvalidPassword
	}

	accessToken, err := jwtauth.GenerateAccessToken(user.ID, []byte(s.cfg.JWT.AccessSecret))
	if err != nil {
		log.Error("generate access token", "error", err)
		return "", "", err
	}

	refreshToken, err := jwtauth.GenerateRefreshToken(user.ID, []byte(s.cfg.JWT.RefreshSecret))
	if err != nil {
		log.Error("generate refresh token", "error", err)
		return "", "", err
	}

	user.JWTRefreshToken = refreshToken
	err = s.repo.SaveUser(*user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
