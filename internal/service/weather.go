package service

import (
	"context"
	"fmt"
	"weather-api/internal/weather"
)

type IWeatherService interface {
	FetchBoth(ctx context.Context, location string) (float64, float64, error)
}

type WeatherService struct {
	weatherApiClient   weather.TemperatureFetcher
	weatherStackClient weather.TemperatureFetcher
}

func NewWeatherService(api weather.TemperatureFetcher, stack weather.TemperatureFetcher) *WeatherService {
	return &WeatherService{
		weatherApiClient:   api,
		weatherStackClient: stack,
	}
}

func (s *WeatherService) FetchBoth(ctx context.Context, location string) (float64, float64, error) {
	type result struct {
		temp float64
		err  error
	}

	ch1 := make(chan result, 1)
	ch2 := make(chan result, 1)

	go func() {

		t, err := s.weatherApiClient.FetchTemperature(ctx, location)
		ch1 <- result{t, err}
	}()

	go func() {

		t, err := s.weatherStackClient.FetchTemperature(ctx, location)
		ch2 <- result{t, err}
	}()

	select {
	case <-ctx.Done():
		return 0, 0, ctx.Err()
	case r1 := <-ch1:
		r2 := <-ch2
		if r1.err != nil || r2.err != nil {
			return 0, 0, fmt.Errorf("one or both services failed")
		}
		return r1.temp, r2.temp, nil
	}
}
