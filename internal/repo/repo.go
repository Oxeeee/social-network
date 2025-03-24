package repo

import (
	"log/slog"

	"gorm.io/gorm"
)

type Repo interface {
}

type repo struct {
	db  *gorm.DB
	log *slog.Logger
}

func NewRepo(db *gorm.DB, log *slog.Logger) Repo {
	return repo{
		db:  db,
		log: log,
	}
}
