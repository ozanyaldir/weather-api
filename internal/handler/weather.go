package handler

import (
	"weather-api/internal/orchestrator"

	"github.com/gofiber/fiber/v2"
)

type IWeatherHandler interface {
	GetWeather(c *fiber.Ctx) error
}

type WeatherHandler struct {
	orchestrator orchestrator.IWeatherOrchestrator
}

func NewWeatherHandler(o orchestrator.IWeatherOrchestrator) *WeatherHandler {
	return &WeatherHandler{
		orchestrator: o,
	}
}

func (h *WeatherHandler) GetWeather(c *fiber.Ctx) error {
	location := c.Query("q")
	if location == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "query parameter 'q' (location) is required",
		})
	}

	summary, err := h.orchestrator.GetWeatherSummary(c.Context(), location)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(summary)
}
