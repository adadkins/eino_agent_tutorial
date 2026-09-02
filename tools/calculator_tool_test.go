package tools

import (
	"context"
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name    string
		input   CalculatorInput
		want    float64
		wantErr bool
	}{
		{"add", CalculatorInput{A: 2, B: 3, Op: "add"}, 5, false},
		{"subtract", CalculatorInput{A: 10, B: 4, Op: "subtract"}, 6, false},
		{"multiply", CalculatorInput{A: 47, B: 6, Op: "multiply"}, 282, false},
		{"divide", CalculatorInput{A: 20, B: 4, Op: "divide"}, 5, false},
		{"divide by zero", CalculatorInput{A: 1, B: 0, Op: "divide"}, 0, true},
		{"unknown op", CalculatorInput{A: 1, B: 1, Op: "modulo"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculator(context.Background(), &tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error state: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
