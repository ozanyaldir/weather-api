package weatherapi

import (
	"context"
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
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) FetchTemperature(ctx context.Context, location string) (float64, error) {
	apiKey := os.Getenv("WEATHERAPI_KEY")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, location)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, fmt.Errorf("weatherapi request failed: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result.Current.TempC, nil
}
