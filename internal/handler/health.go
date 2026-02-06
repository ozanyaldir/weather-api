package handler

import (
	"github.com/gofiber/fiber/v2"
)

type IHealthHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "Weather API is running",
	})
}
