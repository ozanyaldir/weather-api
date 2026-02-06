package route

import (
	"weather-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func registerHealth(app *fiber.App, healthHandler handler.IHealthHandler) {
	app.Get("/health", healthHandler.HealthCheck)
}
