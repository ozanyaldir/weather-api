package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestLoggerMiddleware(t *testing.T) {
	app := fiber.New()
	app.Use(middleware.Logger())

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestErrorHandler(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Get("/fiber-error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusConflict, "already exists")
	})

	app.Get("/generic-error", func(c *fiber.Ctx) error {
		return errors.New("something went wrong")
	})

	t.Run("Handle Fiber Error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/fiber-error", nil)
		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusConflict, resp.StatusCode)
	})

	t.Run("Handle Generic Error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/generic-error", nil)
		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}
