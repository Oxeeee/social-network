package service

import (
	"log/slog"

	"github.com/Oxeeee/social-network/internal/config"
	"github.com/Oxeeee/social-network/internal/models/domain"
	cerrors "github.com/Oxeeee/social-network/internal/models/errors"
	"github.com/Oxeeee/social-network/internal/models/requests"
	"github.com/Oxeeee/social-network/internal/models/responses"
	"github.com/Oxeeee/social-network/internal/repo"
	base64encode "github.com/Oxeeee/social-network/internal/utils/base64"
	"github.com/Oxeeee/social-network/internal/utils/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(req requests.Register) error
	Login(req requests.Login) (*responses.LoginResponse, error)
	LogoutFromAllSessions(userID uint) error
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

	var fileName string
	if req.PhotoEncrypted != "" {
		var err error
		fileName, err = base64encode.FromBase64(req.PhotoEncrypted, req.Username)
		if err != nil {
			log.Error("decode from base 64", "error", err)
			return err
		}
	}

	user := domain.User{
		Email:     req.Email,
		Username:  req.Username,
		Name:      req.Name,
		Surname:   req.Surname,
		PhotoPath: fileName,
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

func (s *service) Login(req requests.Login) (*responses.LoginResponse, error) {
	const op = "service.login"
	log := s.log.With(slog.String("op", op))

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(req.Password)); err != nil {
		log.Debug("compare hash and password", "error", err)
		return nil, cerrors.ErrInvalidPassword
	}

	accessToken, err := jwtauth.GenerateAccessToken(user.ID, []byte(s.cfg.JWT.AccessSecret))
	if err != nil {
		log.Error("generate access token", "error", err)
		return nil, err
	}

	refreshToken, err := jwtauth.GenerateRefreshToken(user.ID, user.RefreshTokenVersion, []byte(s.cfg.JWT.RefreshSecret))
	if err != nil {
		log.Error("generate refresh token", "error", err)
		return nil, err
	}

	var photoEncoded string
	if user.PhotoPath != "" {
		var err error
		photoEncoded, err = base64encode.ToBase64(user.PhotoPath)
		if err != nil {
			log.Error("encode to base64", "error", err)
			return nil, err
		}
	}
	

	resp := responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Username:     user.Username,
		Name:         user.Name,
		Surname:      user.Surname,
	}
	
	if photoEncoded != "" {
		resp.Photo = photoEncoded
	}

	return &resp, nil
}

func (s *service) LogoutFromAllSessions(userID uint) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.RefreshTokenVersion += 1

	if err := s.repo.SaveUser(*user); err != nil {
		return err
	}
	return nil
}
