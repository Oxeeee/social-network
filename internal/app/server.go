package app

import (
	"log/slog"
	"time"

	"github.com/Oxeeee/shopping-yona/internal/transport/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	engine *gin.Engine
}

func New(log *slog.Logger, handlers handlers.Handlers) *App {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://theca.oxytocingroup.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	engine.SetTrustedProxies(nil)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return &App{engine: engine}
}

func (a *App) Start() {
	a.engine.Run(":3000")
}
