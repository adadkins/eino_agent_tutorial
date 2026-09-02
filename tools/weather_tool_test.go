package tools

import (
	"context"
	"testing"
)

func TestWeather(t *testing.T) {
	out, err := Weather(context.Background(), &WeatherInput{City: "London"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.City == "" {
		t.Error("expected a resolved city name")
	}
}

func TestWeather_UnknownCity(t *testing.T) {
	_, err := Weather(context.Background(), &WeatherInput{City: "Not A Real City XYZ123"})
	if err == nil {
		t.Error("expected an error for an unknown city")
	}
}
