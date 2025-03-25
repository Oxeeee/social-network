package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	cerrors "github.com/Oxeeee/social-network/internal/models/errors"
	"github.com/Oxeeee/social-network/internal/models/requests"
	"github.com/Oxeeee/social-network/internal/models/responses"
	"github.com/Oxeeee/social-network/internal/service"
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	LogoutFromAllSessions(c echo.Context) error
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
			c.JSON(http.StatusBadRequest, responses.Response{Error: err.Error()})
		} else if errors.Is(err, cerrors.ErrEmailTaken) {
			c.JSON(http.StatusBadRequest, responses.Response{Error: err.Error()})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Can not register user: %s", err))
	}

	return c.JSON(http.StatusCreated, responses.Response{Message: "user registred sucessfully"})
}

func (h *handlers) Login(c echo.Context) error {
	var req requests.Login
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	accessToken, refreshToken, err := h.service.Login(req)
	if err != nil {
		if errors.Is(err, cerrors.ErrInvalidEmail) {
			return c.JSON(http.StatusBadRequest, responses.Response{Error: err.Error()})
		}
		if errors.Is(err, cerrors.ErrInvalidPassword) {
			return c.JSON(http.StatusUnauthorized, responses.Response{Error: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responses.Response{Error: err.Error()})
	}

	c.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return c.JSON(http.StatusOK, responses.Response{Message: "user logged in successfully", Details: map[string]any{"accessToken": accessToken}})
}

func (h *handlers) Logout(c echo.Context) error {
	h.log.Debug(c.Get("userID").(string))
	c.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return c.JSON(http.StatusOK, responses.Response{Message: "user logged out successfully"})
}

func (h *handlers) LogoutFromAllSessions(c echo.Context) error {
	const op = "handlers.logoutFromAllSessions"
	log := h.log.With(slog.String("op", op))
	userID := c.Get("userID")
	if userID == nil {
		log.Error("not found userID in context")
		return c.JSON(http.StatusInternalServerError, responses.Response{Error: "didnt find userID in context value"})
	}

	err := h.service.LogoutFromAllSessions(userID.(uint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{Error: err.Error()})
	}

	c.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return c.JSON(http.StatusOK, responses.Response{Message: "logged out from all sessions"})
}
