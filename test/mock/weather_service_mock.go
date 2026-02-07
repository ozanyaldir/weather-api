package mock

import (
	"context"
	"fmt"
	"sync/atomic"
	"weather-api/internal/weather"
)

type MockWeatherService struct {
	CallCount int32
	Temp1     float64
	Temp2     float64
	Err       error
}

func (m *MockWeatherService) FetchBoth(ctx context.Context, location string) (float64, float64, error) {
	atomic.AddInt32(&m.CallCount, 1)
	return m.Temp1, m.Temp2, m.Err
}

type MockFetcher struct {
	Temp       float64
	ShouldFail bool
}

var _ weather.TemperatureFetcher = (*MockFetcher)(nil)

func (m *MockFetcher) FetchTemperature(ctx context.Context, city string) (float64, error) {
	if m.ShouldFail {
		return 0, fmt.Errorf("simulated api error for city: %s", city)
	}
	return m.Temp, nil
}
