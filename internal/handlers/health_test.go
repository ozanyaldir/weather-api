package handlers

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHealthCheck(t *testing.T) {
	app := fiber.New()
	app.Get("/health", HealthCheck)

	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Failed to test request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if len(body) == 0 {
		t.Error("Response body should not be empty")
	}
}

func TestNewHealthHandler(t *testing.T) {
	handler := NewHealthHandler()

	if handler == nil {
		t.Fatal("Handler should not be nil")
	}
}
