package service

import (
	"context"
	"sync"
	"time"
	"weather-api/internal/model"
)

type IWeatherBatchService interface {
	GetWeather(ctx context.Context, location string) (float64, float64, int, error)
}

type WeatherBatchService struct {
	weatherService IWeatherService
	mu             sync.Mutex
	batches        map[string]*batch
}

type batch struct {
	location string
	chans    []chan model.BatchResult
	timer    *time.Timer
}

func NewWeatherBatchService(ws IWeatherService) *WeatherBatchService {
	return &WeatherBatchService{
		weatherService: ws,
		batches:        make(map[string]*batch),
	}
}

func (s *WeatherBatchService) GetWeather(ctx context.Context, location string) (float64, float64, int, error) {
	b := s.getOrCreateBatch(location)
	resChan := s.joinBatch(b)

	select {
	case res := <-resChan:
		return res.Temp1, res.Temp2, res.Count, res.Err
	case <-ctx.Done():
		return 0, 0, 0, ctx.Err()
	}
}

func (s *WeatherBatchService) getOrCreateBatch(location string) *batch {
	s.mu.Lock()
	defer s.mu.Unlock()

	if b, exists := s.batches[location]; exists {
		return b
	}

	b := &batch{
		location: location,
		chans:    make([]chan model.BatchResult, 0, 10),
	}

	b.timer = time.AfterFunc(5*time.Second, func() {
		s.executeBatch(location)
	})

	s.batches[location] = b
	return b
}

func (s *WeatherBatchService) joinBatch(b *batch) chan model.BatchResult {
	s.mu.Lock()
	defer s.mu.Unlock()

	resChan := make(chan model.BatchResult, 1)
	b.chans = append(b.chans, resChan)

	if len(b.chans) == 10 {
		if b.timer.Stop() {
			go s.executeBatch(b.location)
		}
	}

	return resChan
}

func (s *WeatherBatchService) executeBatch(location string) {
	s.mu.Lock()
	b, exists := s.batches[location]
	if !exists {
		s.mu.Unlock()
		return
	}
	delete(s.batches, location)
	s.mu.Unlock()

	t1, t2, err := s.weatherService.FetchBoth(context.Background(), b.location)

	res := model.BatchResult{
		Temp1: t1,
		Temp2: t2,
		Count: len(b.chans),
		Err:   err,
	}

	for _, ch := range b.chans {
		ch <- res
	}
}
