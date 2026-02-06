package weatherstack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"weather-api/internal/weather"
)

type Client struct {
	httpClient *http.Client
}

func New() weather.TemperatureFetcher {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) FetchTemperature(location string) (float64, error) {
	apiKey := os.Getenv("WEATHERSTACK_KEY")

	url := fmt.Sprintf(
		"http://api.weatherstack.com/current?access_key=%s&query=%s",
		apiKey,
		location,
	)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, fmt.Errorf("weatherstack request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("weatherstack returned status %d", resp.StatusCode)
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("weatherstack decode failed: %w", err)
	}

	return result.Current.Temperature, nil
}
