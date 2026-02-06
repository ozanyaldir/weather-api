package app

import (
	"weather-api/internal/handler"
	"weather-api/internal/orchestrator"
	"weather-api/internal/pkg/weatherapi"
	"weather-api/internal/pkg/weatherstack"
	"weather-api/internal/repository"
	"weather-api/internal/route"
	"weather-api/internal/service"
	"weather-api/internal/weather"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Config struct {
	Fiber   fiber.Config
	DB      *gorm.DB
	Weather weather.TemperatureFetcher
	Stack   weather.TemperatureFetcher
}

func Bootstrap(cfg Config) *fiber.App {
	wAPI := cfg.Weather
	if wAPI == nil {
		wAPI = weatherapi.New()
	}

	wStack := cfg.Stack
	if wStack == nil {
		wStack = weatherstack.New()
	}

	wRepo := repository.NewWeatherQueryRepository(cfg.DB)
	wService := service.NewWeatherService(wAPI, wStack)
	wBatch := service.NewWeatherBatchService(wService)
	wOrch := orchestrator.NewWeatherOrchestrator(wBatch, wRepo)

	wHandler := handler.NewWeatherHandler(wOrch)
	hHandler := handler.NewHealthHandler()

	app := fiber.New(cfg.Fiber)
	route.Register(app, wHandler, hHandler)

	return app
}
