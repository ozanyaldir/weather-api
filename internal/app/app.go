package app

import (
	"weather-api/internal/database"
	"weather-api/internal/handler"
	"weather-api/internal/orchestrator"
	"weather-api/internal/pkg/weatherapi"
	"weather-api/internal/pkg/weatherstack"
	"weather-api/internal/repository"
	"weather-api/internal/route"
	"weather-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Bootstrap(fiberCfg fiber.Config) *fiber.App {
	wAPI := weatherapi.New()
	wStack := weatherstack.New()

	wRepo := repository.NewWeatherQueryRepository(database.DB)

	wService := service.NewWeatherService(wAPI, wStack)

	wOrch := orchestrator.NewWeatherOrchestrator(wService, wRepo)

	wHandler := handler.NewWeatherHandler(wOrch)
	hHandler := handler.NewHealthHandler()

	app := fiber.New(fiberCfg)
	route.Register(app, wHandler, hHandler)

	return app
}
