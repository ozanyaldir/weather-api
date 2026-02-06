package route

import (
	"weather-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func registerWeather(app *fiber.App, weatherHandler handler.IWeatherHandler) {
	app.Get("/weather", weatherHandler.GetWeather)
}
