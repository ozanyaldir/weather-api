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
	batchService service.IWeatherBatchService
	repo         repository.IWeatherQueryRepository
}

func NewWeatherOrchestrator(bs service.IWeatherBatchService, r repository.IWeatherQueryRepository) *WeatherOrchestrator {
	return &WeatherOrchestrator{
		batchService: bs,
		repo:         r,
	}
}

func (o *WeatherOrchestrator) GetWeatherSummary(ctx context.Context, location string) (dto.WeatherResponse, error) {
	temp1, temp2, count, err := o.batchService.GetWeather(ctx, location)
	if err != nil {
		return dto.WeatherResponse{}, fmt.Errorf("orchestrator batch fetch failed: %w", err)
	}

	avg := (temp1 + temp2) / 2

	go o.repo.Create(location, temp1, temp2, count)

	return dto.WeatherResponse{
		Location:    location,
		Temperature: avg,
	}, nil
}
