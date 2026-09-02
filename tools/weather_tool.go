package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Weather(ctx context.Context, input *WeatherInput) (*WeatherOutput, error) {
	// Step 1: geocode city name to lat/lon
	geoURL := fmt.Sprintf(
		"https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1",
		url.QueryEscape(input.City),
	)
	geoReq, err := http.NewRequestWithContext(ctx, http.MethodGet, geoURL, nil)
	if err != nil {
		return nil, err
	}
	geoResp, err := http.DefaultClient.Do(geoReq)
	if err != nil {
		return nil, fmt.Errorf("geocoding request failed: %w", err)
	}
	defer geoResp.Body.Close()
	geoBody, err := io.ReadAll(geoResp.Body)
	if err != nil {
		return nil, err
	}

	var geo struct {
		Results []struct {
			Name      string  `json:"name"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"results"`
	}
	if err := json.Unmarshal(geoBody, &geo); err != nil {
		return nil, fmt.Errorf("failed to parse geocoding response: %w", err)
	}
	if len(geo.Results) == 0 {
		return nil, fmt.Errorf("could not find city: %s", input.City)
	}
	lat, lon := geo.Results[0].Latitude, geo.Results[0].Longitude

	// Step 2: fetch current weather for that location
	fcURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", lat, lon)
	fcReq, err := http.NewRequestWithContext(ctx, http.MethodGet, fcURL, nil)
	if err != nil {
		return nil, err
	}
	fcResp, err := http.DefaultClient.Do(fcReq)
	if err != nil {
		return nil, fmt.Errorf("forecast request failed: %w", err)
	}
	defer fcResp.Body.Close()
	fcBody, err := io.ReadAll(fcResp.Body)
	if err != nil {
		return nil, err
	}

	var fc struct {
		CurrentWeather struct {
			Temperature float64 `json:"temperature"`
			WindSpeed   float64 `json:"windspeed"`
		} `json:"current_weather"`
	}
	if err := json.Unmarshal(fcBody, &fc); err != nil {
		return nil, fmt.Errorf("failed to parse forecast response: %w", err)
	}

	return &WeatherOutput{
		City:        geo.Results[0].Name,
		TempC:       fc.CurrentWeather.Temperature,
		WindKph:     fc.CurrentWeather.WindSpeed,
		Description: fmt.Sprintf("%.1f°C, wind %.1f km/h", fc.CurrentWeather.Temperature, fc.CurrentWeather.WindSpeed),
	}, nil
}
