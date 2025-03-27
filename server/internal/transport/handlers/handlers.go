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

// Register godoc
// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя в системе
// @Tags auth
// @Accept json
// @Produce json
// @Param request body requests.Register true "Данные для регистрации"
// @Success 201 {object} responses.Response "Пользователь успешно зарегистрирован"
// @Failure 400 {object} responses.Response "Ошибка валидации | USERNAME_ALREADY_TAKEN | EMAIL_ALREADY_TAKEN"
// @Failure 500 {object} responses.Response "Внутренняя ошибка сервера"
// @Router /register [post]
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
			c.JSON(http.StatusBadRequest, responses.Response[any]{Error: err.Error()})
		} else if errors.Is(err, cerrors.ErrEmailTaken) {
			c.JSON(http.StatusBadRequest, responses.Response[any]{Error: err.Error()})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Can not register user: %s", err))
	}

	return c.JSON(http.StatusCreated, responses.Response[any]{Message: "user registred sucessfully"})
}

// Login godoc
// @Summary Вход пользователя
// @Description Авторизует пользователя и выдает токены доступа
// @Tags auth
// @Accept json
// @Produce json
// @Param request body requests.Login true "Данные для входа"
// @Success 200 {object} responses.Response "Пользователь успешно авторизован, возвращает accessToken"
// @Failure 400 {object} responses.Response "Ошибка валидации | INVALID_EMAIL"
// @Failure 401 {object} responses.Response "INVALID_PASSWORD"
// @Failure 500 {object} responses.Response "Внутренняя ошибка сервера"
// @Router /login [post]
func (h *handlers) Login(c echo.Context) error {
	var req requests.Login
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loginResp, err := h.service.Login(req)
	if err != nil {
		if errors.Is(err, cerrors.ErrInvalidEmail) {
			return c.JSON(http.StatusBadRequest, responses.Response[any]{Error: err.Error()})
		}
		if errors.Is(err, cerrors.ErrInvalidPassword) {
			return c.JSON(http.StatusUnauthorized, responses.Response[any]{Error: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responses.Response[any]{Error: err.Error()})
	}

	c.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    loginResp.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	resp := responses.Response[responses.LoginResponse]{
		Data:    *loginResp,
		Message: "user logged in successfully",
	}

	return c.JSON(http.StatusOK, resp)
}

// Logout godoc
// @Summary Выход пользователя
// @Description Выполняет выход пользователя из текущей сессии
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.Response "Пользователь успешно вышел"
// @Failure 401 {object} responses.Response "MISSING_AUTHORIZATION_TOKEN | INVALID_AUTHORIZATION_HEADER_FORMAT | INVALID_OR_EXPIRED_TOKEN"
// @Failure 500 {object} responses.Response "Внутренняя ошибка сервера"
// @Router /auth/logout [post]
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
	return c.JSON(http.StatusOK, responses.Response[any]{Message: "user logged out successfully"})
}

// LogoutFromAllSessions godoc
// @Summary Выход из всех сессий
// @Description Выполняет выход пользователя из всех активных сессий
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.Response "Выход из всех сессий выполнен успешно"
// @Failure 401 {object} responses.Response "MISSING_AUTHORIZATION_TOKEN | INVALID_AUTHORIZATION_HEADER_FORMAT | INVALID_OR_EXPIRED_TOKEN"
// @Failure 500 {object} responses.Response "Внутренняя ошибка сервера | Не найден userID в контексте"
// @Router /auth/logout/all [post]
func (h *handlers) LogoutFromAllSessions(c echo.Context) error {
	const op = "handlers.logoutFromAllSessions"
	log := h.log.With(slog.String("op", op))
	userID := c.Get("userID")
	if userID == nil {
		log.Error("not found userID in context")
		return c.JSON(http.StatusInternalServerError, responses.Response[any]{Error: "didnt find userID in context value"})
	}

	err := h.service.LogoutFromAllSessions(userID.(uint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response[any]{Error: err.Error()})
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

	return c.JSON(http.StatusOK, responses.Response[any]{Message: "logged out from all sessions"})
}
