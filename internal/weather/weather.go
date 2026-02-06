package weather

import "context"

type TemperatureFetcher interface {
	FetchTemperature(ctx context.Context, location string) (float64, error)
}
