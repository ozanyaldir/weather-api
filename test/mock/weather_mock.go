package mock

import "context"

type MockWeatherFetcher struct {
	Temp float64
	Err  error
}

func (m *MockWeatherFetcher) FetchTemperature(ctx context.Context, location string) (float64, error) {
	return m.Temp, m.Err
}
