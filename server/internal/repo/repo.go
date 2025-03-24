package repo

import (
	"log/slog"

	"github.com/Oxeeee/social-network/internal/models/domain"
	cerrors "github.com/Oxeeee/social-network/internal/models/errors"
	"gorm.io/gorm"
)

type Repo interface {
	Register(user domain.User) error
	GetUserByUsername(username string) (*domain.User, error)
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

func (r *repo) GetUserByUsername(username string) (*domain.User, error) {
	const op = "repo.getUserByUsername"
	log := r.log.With(slog.String("op", op))
	var user domain.User
	err := r.db.Model(&domain.User{}).Where("username = ?", username).Find(&user).Error
	if err != nil {
		log.Error("get by username", "error", err)
		return nil, err
	}

	return &user, nil
}
