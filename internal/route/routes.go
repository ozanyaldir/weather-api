package route

import (
	"weather-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, weatherHandler handler.IWeatherHandler, healthHandler handler.IHealthHandler) {
	registerHealth(app, healthHandler)
	registerWeather(app, weatherHandler)
}
