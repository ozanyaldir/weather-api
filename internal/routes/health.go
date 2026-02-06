package routes

import (
	"weather-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func registerHealth(app *fiber.App) {
	app.Get("/health", handlers.HealthCheck)
}
