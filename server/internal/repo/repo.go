package repo

import (
	"log/slog"

	"github.com/Oxeeee/social-network/internal/models/domain"
	cerrors "github.com/Oxeeee/social-network/internal/models/errors"
	"gorm.io/gorm"
)

type Repo interface {
	Register(user domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	SaveUser(user domain.User) error
}

type repo struct {
	db  *gorm.DB
	log *slog.Logger
}

func NewRepo(db *gorm.DB, log *slog.Logger) Repo {
	return &repo{
		db:  db,
		log: log,
	}
}

func (r *repo) Register(user domain.User) error {
	const op = "repo.register"
	log := r.log.With(slog.String("op", op))

	var c int64
	err := r.db.Model(&domain.User{}).Where("username = ?", user.Username).Count(&c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error("get usename", "error", err)
		return err
	}
	if c > 0 {
		return cerrors.ErrUsernameTaken
	}

	err = r.db.Model(&domain.User{}).Where("email = ?", user.Email).Count(&c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error("get email", "error", err)
		return err
	}
	if c > 0 {
		return cerrors.ErrEmailTaken
	}

	err = r.db.Model(&domain.User{}).Create(&user).Error
	if err != nil {
		log.Error("create", "error", err)
		return err
	}

	return nil
}

func (r *repo) GetUserByEmail(email string) (*domain.User, error) {
	const op = "repo.getUserByEmail"
	log := r.log.With(slog.String("op", op))
	var user domain.User
	if err := r.db.Model(&domain.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Debug("get by email", "error", err)
			return nil, cerrors.ErrInvalidEmail
		}
		log.Error("get by email", "error", err)
		return nil, err
	}

	return &user, nil
}

func (r *repo) SaveUser(user domain.User) error {
	const op = "repo.saveUser"
	log := r.log.With(slog.String("op", op))
	err := r.db.Model(&domain.User{}).Where("id = ?", user.ID).Save(user).Error
	if err != nil {
		log.Error("save user", "error", err)
		return err
	}

	return nil
}
