package handler_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
	"weather-api/internal/dto"
	"weather-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type MockOrchestrator struct {
	Response dto.WeatherResponse
	Err      error
}

func (m *MockOrchestrator) GetWeatherSummary(ctx context.Context, location string) (dto.WeatherResponse, error) {
	return m.Response, m.Err
}

func TestGetWeather(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		app := fiber.New()

		mockData := dto.WeatherResponse{
			Location:    "istanbul",
			Temperature: 25.0,
		}
		mockOrch := &MockOrchestrator{Response: mockData}

		h := handler.NewWeatherHandler(mockOrch)
		app.Get("/weather", h.GetWeather)

		req := httptest.NewRequest("GET", "/weather?q=istanbul", nil)
		resp, err := app.Test(req, -1)

		if err != nil {
			t.Fatalf("Failed to test request: %v", err)
		}

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		var responseData dto.WeatherResponse
		if err := json.Unmarshal(body, &responseData); err != nil {
			t.Fatalf("Failed to decode JSON: %v", err)
		}

		if responseData.Location != "istanbul" {
			t.Errorf("Expected location 'istanbul', got '%s'", responseData.Location)
		}
	})

	t.Run("MissingQueryParam", func(t *testing.T) {
		app := fiber.New()
		mockOrch := &MockOrchestrator{}
		h := handler.NewWeatherHandler(mockOrch)
		app.Get("/weather", h.GetWeather)

		req := httptest.NewRequest("GET", "/weather", nil)
		resp, _ := app.Test(req, -1)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", resp.StatusCode)
		}
	})

	t.Run("OrchestratorError", func(t *testing.T) {
		app := fiber.New()
		mockOrch := &MockOrchestrator{
			Err: fmt.Errorf("api failure"),
		}
		h := handler.NewWeatherHandler(mockOrch)
		app.Get("/weather", h.GetWeather)

		req := httptest.NewRequest("GET", "/weather?q=istanbul", nil)
		resp, _ := app.Test(req, -1)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("Expected status 500, got %d", resp.StatusCode)
		}
	})
}

func TestNewWeatherHandler(t *testing.T) {
	mockOrch := &MockOrchestrator{}
	h := handler.NewWeatherHandler(mockOrch)

	if h == nil {
		t.Fatal("WeatherHandler should not be nil")
	}
}
