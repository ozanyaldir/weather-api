package service

import (
	"fmt"
	"weather-api/internal/weather"
)

type IWeatherService interface {
	FetchBoth(location string) (float64, float64, error)
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

func (s *WeatherService) FetchBoth(location string) (float64, float64, error) {
	type result struct {
		temp float64
		err  error
	}

	ch1 := make(chan result, 1)
	ch2 := make(chan result, 1)

	go func() {
		t, err := s.weatherApiClient.FetchTemperature(location)
		ch1 <- result{t, err}
	}()

	go func() {
		t, err := s.weatherStackClient.FetchTemperature(location)
		ch2 <- result{t, err}
	}()

	r1 := <-ch1
	r2 := <-ch2

	if r1.err != nil || r2.err != nil {
		return 0, 0, fmt.Errorf("one or both services failed")
	}

	return r1.temp, r2.temp, nil
}
