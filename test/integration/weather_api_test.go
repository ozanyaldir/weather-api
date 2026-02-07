package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"weather-api/internal/app"
	"weather-api/internal/dto"
	"weather-api/test/mock"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestWeatherAPI_Integration(t *testing.T) {

	mockDB, sqlMock, _ := mock.NewGormMock()

	mockAPI := &mock.MockWeatherFetcher{Temp: 10.0}
	mockStack := &mock.MockWeatherFetcher{Temp: 20.0}

	server := app.Bootstrap(app.Config{
		Fiber:   fiber.Config{AppName: "Integration Test"},
		DB:      mockDB,
		Weather: mockAPI,
		Stack:   mockStack,
	})

	t.Run("GET /weather - Success Flow", func(t *testing.T) {

		sqlMock.ExpectBegin()
		sqlMock.ExpectExec("INSERT INTO `weather_queries`").
			WillReturnResult(mock.NewResult(1, 1))
		sqlMock.ExpectCommit()

		req := httptest.NewRequest("GET", "/weather?q=Istanbul", nil)
		resp, err := server.Test(req, 10000)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var body dto.WeatherResponse
		json.NewDecoder(resp.Body).Decode(&body)

		assert.Equal(t, "Istanbul", body.Location)
		assert.Equal(t, 15.0, body.Temperature)

		time.Sleep(100 * time.Millisecond)

		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("GET /health - Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/health", nil)
		resp, err := server.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var body map[string]string
		json.NewDecoder(resp.Body).Decode(&body)
		assert.Equal(t, "ok", body["status"])
	})
}
