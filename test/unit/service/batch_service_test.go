package service_test

import (
	"context"
	"testing"
	"weather-api/internal/service"
	"weather-api/test/mock"
)

func TestBatchService_WithSharedMock(t *testing.T) {
	mockWS := &mock.MockWeatherService{
		Temp1: 22.5,
		Temp2: 25.0,
	}

	svc := service.NewWeatherBatchService(mockWS)

	_, _, _, err := svc.GetWeather(context.Background(), "Istanbul")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if mockWS.CallCount != 1 {
		t.Errorf("Expected 1 call, got %d", mockWS.CallCount)
	}
}
