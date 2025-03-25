package loggermiddleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		// Обработка запроса
		err := next(c)

		// Логируем после обработки
		stop := time.Now()
		latency := stop.Sub(start)

		method := c.Request().Method
		uri := c.Request().RequestURI
		status := c.Response().Status
		ip := c.RealIP()

		fmt.Printf("[%-7s] %3d | %-15s | %-30s | %s\n",
			method,
			status,
			ip,
			uri,
			latency,
		)

		return err
	}
}
