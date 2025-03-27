package app

import (
	"log/slog"
	"time"

	_ "github.com/Oxeeee/social-network/docs"
	"github.com/Oxeeee/social-network/internal/transport/handlers"
	authmw "github.com/Oxeeee/social-network/internal/utils/authmiddleware"
	customvalidator "github.com/Oxeeee/social-network/internal/utils/validator"
	loggermiddleware "github.com/Oxeeee/social-network/pkg/logger_middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type App struct {
	engine *echo.Echo
}

func New(log *slog.Logger, handlers handlers.Handlers, mw authmw.AuthMiddleware) *App {
	e := echo.New()
	e.Validator = &customvalidator.CustomValidator{Validator: validator.New()}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour / time.Second),
	}))
	e.Use(loggermiddleware.RequestLogger)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	secure := e.Group("/secure", mw.JWTMiddleware)
	secure.POST("/logout", handlers.Logout)
	secure.POST("/logout/all", handlers.LogoutFromAllSessions)

	return &App{engine: e}
}

func (a *App) Start() {
	a.engine.Start(":3000")
}
