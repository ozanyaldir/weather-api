package orchestrator

import (
	"context"
	"fmt"
	"weather-api/internal/dto"
	"weather-api/internal/repository"
	"weather-api/internal/service"
)

type IWeatherOrchestrator interface {
	GetWeatherSummary(ctx context.Context, location string) (dto.WeatherResponse, error)
}

type WeatherOrchestrator struct {
	service service.IWeatherService
	repo    repository.IWeatherQueryRepository
}

func NewWeatherOrchestrator(s service.IWeatherService, r repository.IWeatherQueryRepository) *WeatherOrchestrator {
	return &WeatherOrchestrator{
		service: s,
		repo:    r,
	}
}

func (o *WeatherOrchestrator) GetWeatherSummary(ctx context.Context, location string) (dto.WeatherResponse, error) {
	temp1, temp2, err := o.service.FetchBoth(location)
	if err != nil {
		return dto.WeatherResponse{}, fmt.Errorf("orchestrator fetch failed: %w", err)
	}

	avg := (temp1 + temp2) / 2

	go o.repo.Create(location, temp1, temp2, 1)

	return dto.WeatherResponse{
		Location:    location,
		Temperature: avg,
	}, nil
}
