package weather

type TemperatureFetcher interface {
	FetchTemperature(city string) (float64, error)
}
