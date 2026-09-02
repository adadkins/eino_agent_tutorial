package tools

import (
	"context"
	"time"
)

func MoonPhase(ctx context.Context, input *MoonPhaseInput) (*MoonPhaseOutput, error) {
	// Known reference new moon: Jan 6, 2000, 18:14 UTC
	knownNewMoon := time.Date(2000, 1, 6, 18, 14, 0, 0, time.UTC)
	synodicMonth := 29.53058867 // average days per lunar cycle

	daysSince := time.Since(knownNewMoon).Hours() / 24
	age := daysSince - float64(int(daysSince/synodicMonth))*synodicMonth
	if age < 0 {
		age += synodicMonth
	}

	var phase string
	switch {
	case age < 1.84:
		phase = "New Moon"
	case age < 5.53:
		phase = "Waxing Crescent"
	case age < 9.22:
		phase = "First Quarter"
	case age < 12.91:
		phase = "Waxing Gibbous"
	case age < 16.61:
		phase = "Full Moon"
	case age < 20.30:
		phase = "Waning Gibbous"
	case age < 23.99:
		phase = "Last Quarter"
	case age < 27.68:
		phase = "Waning Crescent"
	default:
		phase = "New Moon"
	}

	city := input.City
	if city == "" {
		city = "Los Angeles"
	}

	return &MoonPhaseOutput{City: city, Phase: phase, Age: age}, nil
}
