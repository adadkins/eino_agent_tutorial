package tools

import (
	"context"
	"testing"
)

func TestMoonPhase(t *testing.T) {
	out, err := MoonPhase(context.Background(), &MoonPhaseInput{City: "London"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.City != "London" {
		t.Errorf("expected city London, got %s", out.City)
	}
	if out.Age < 0 || out.Age > 29.6 {
		t.Errorf("age out of valid range: %v", out.Age)
	}
	if out.Phase == "" {
		t.Error("expected a non-empty phase")
	}
}

func TestMoonPhase_DefaultsCity(t *testing.T) {
	out, err := MoonPhase(context.Background(), &MoonPhaseInput{City: ""})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.City != "Los Angeles" {
		t.Errorf("expected default city Los Angeles, got %s", out.City)
	}
}
