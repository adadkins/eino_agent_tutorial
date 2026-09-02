package tools

import (
	"context"
	"fmt"
)

func Calculator(ctx context.Context, input *CalculatorInput) (float64, error) {
	switch input.Op {
	case "add":
		return input.A + input.B, nil
	case "subtract":
		return input.A - input.B, nil
	case "multiply":
		return input.A * input.B, nil
	case "divide":
		if input.B == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return input.A / input.B, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", input.Op)
	}
}
