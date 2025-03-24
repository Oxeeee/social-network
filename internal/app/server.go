package app

import (
	"log/slog"
	"time"

	"github.com/Oxeeee/social-network/internal/transport/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type App struct {
	engine *echo.Echo
}

func New(log *slog.Logger, handlers handlers.Handlers) *App {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173", "https://theca.oxytocingroup.com"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour / time.Second),
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return &App{engine: e}
}

func (a *App) Start() {
	a.engine.Start(":3000")
}
