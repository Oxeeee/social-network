package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	cerrors "github.com/Oxeeee/social-network/internal/models/errors"
	"github.com/Oxeeee/social-network/internal/models/requests"
	"github.com/Oxeeee/social-network/internal/models/responses"
	"github.com/Oxeeee/social-network/internal/service"
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	HelloWorld(c echo.Context) error
	Register(c echo.Context) error
}

type handlers struct {
	log     *slog.Logger
	service service.Service
}

func NewHandler(log *slog.Logger, service service.Service) Handlers {
	return &handlers{
		log:     log,
		service: service,
	}
}

func (h *handlers) HelloWorld(c echo.Context) error {
	c.JSON(200, "Hello mfker")
	return nil
}

func (h *handlers) Register(c echo.Context) error {
	var req requests.Register
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.service.Register(req)
	if err != nil {
		if errors.Is(err, cerrors.ErrUsernameTaken) {
			c.JSON(http.StatusBadRequest, responses.Response{Error: err})
		} else if errors.Is(err, cerrors.ErrEmailTaken) {
			c.JSON(http.StatusBadRequest, responses.Response{Error: err})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Can not register user: %s", err))
	}

	return c.JSON(http.StatusOK, responses.Response{Message: "user registred sucessfully"})
}
