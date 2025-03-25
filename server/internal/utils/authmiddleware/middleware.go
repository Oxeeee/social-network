package authmw

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/Oxeeee/social-network/internal/config"
	cerrors "github.com/Oxeeee/social-network/internal/models/errors"
	"github.com/Oxeeee/social-network/internal/models/responses"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware interface {
	JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type middleware struct {
	log *slog.Logger
	cfg *config.Config
}

func NewAuthMiddleware(log *slog.Logger, cfg *config.Config) AuthMiddleware {
	return &middleware{
		log: log,
		cfg: cfg,
	}
}

func (mw *middleware) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "mw.JWTMiddleware"
		log := mw.log.With(slog.String("op", op))
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, responses.Response{Error: cerrors.ErrMissingToken.Error()})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.JSON(http.StatusUnauthorized, responses.Response{Error: cerrors.ErrInvalidAuthHeaderFormat.Error()})
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, c.JSON(http.StatusUnauthorized, responses.Response{Error: cerrors.ErrUnexpectedSigningMethod.Error()})
			}
			return []byte(mw.cfg.JWT.AccessSecret), nil
		})

		if err != nil || !token.Valid {
			log.Debug("error", "error", err, "tokenStr", tokenStr)
			return c.JSON(http.StatusUnauthorized, responses.Response{Error: cerrors.ErrInvalidExpToken.Error()})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
				return c.JSON(http.StatusUnauthorized, cerrors.ErrExpToken)
			}
			if userIDFloat, ok := claims["userId"].(float64); ok {
				userID := uint(userIDFloat)
				c.Set("userID", userID)
			} else {
				log.Debug("error INV PAY")
				return c.JSON(http.StatusUnauthorized, responses.Response{Error: cerrors.ErrInvalidPayload.Error()})
			}
		} else {
			log.Debug("error INV PAY 2")
			return c.JSON(http.StatusUnauthorized, responses.Response{Error: cerrors.ErrInvalidPayload.Error()})
		}
		return next(c)
	}
}
