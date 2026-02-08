package handler_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"
	"weather-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func TestHealthCheck(t *testing.T) {
	app := fiber.New()
	handler := handler.NewHealthHandler()

	app.Get("/health", handler.HealthCheck)

	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req, -1)

	if err != nil {
		t.Fatalf("Failed to test request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	var responseData map[string]string
	if err := json.Unmarshal(body, &responseData); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	if responseData["status"] != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", responseData["status"])
	}
}

func TestNewHealthHandler(t *testing.T) {
	handler := handler.NewHealthHandler()

	if handler == nil {
		t.Fatal("Handler should not be nil")
	}
}
