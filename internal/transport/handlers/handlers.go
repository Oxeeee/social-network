package handlers

import (
	"log/slog"

	"github.com/Oxeeee/social-network/internal/service"
)

type Handlers interface {
}

type handlers struct {
	log     *slog.Logger
	service service.Service
}

func NewHandler(log *slog.Logger, service service.Service) Handlers {
	return handlers{
		log:     log,
		service: service,
	}
}
