package service_test

import (
	"context"
	"testing"
	"weather-api/internal/service"
	"weather-api/test/mock"
)

func TestWeatherService_FetchBoth_ParallelSuccess(t *testing.T) {
	apiClient := &mock.MockFetcher{Temp: 18.5}
	stackClient := &mock.MockFetcher{Temp: 21.0}

	svc := service.NewWeatherService(apiClient, stackClient)

	t1, t2, err := svc.FetchBoth(context.Background(), "istanbul")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if t1 != 18.5 {
		t.Errorf("Expected 18.5 from API Client, got %v", t1)
	}
	if t2 != 21.0 {
		t.Errorf("Expected 21.0 from WeatherStack, got %v", t2)
	}
}

func TestWeatherService_FetchBoth_OneFails(t *testing.T) {
	apiClient := &mock.MockFetcher{ShouldFail: true}
	stackClient := &mock.MockFetcher{Temp: 20.0}

	svc := service.NewWeatherService(apiClient, stackClient)

	_, _, err := svc.FetchBoth(context.Background(), "istanbul")

	if err == nil {
		t.Fatal("Expected an error because one client failed, but got nil")
	}

	expectedErr := "one or both services failed"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message '%s', got '%s'", expectedErr, err.Error())
	}
}
