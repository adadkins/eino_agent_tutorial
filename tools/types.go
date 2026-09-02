package tools

// API request/response types
type ChatRequest struct {
	Prompt string `json:"prompt"`
}

type ChatResponse struct {
	Answer   string `json:"answer"`
	ToolUsed bool   `json:"tool_used"`
}

// Calculator
type CalculatorInput struct {
	A  float64 `json:"a" jsonschema_description:"The first number"`
	B  float64 `json:"b" jsonschema_description:"The second number"`
	Op string  `json:"op" jsonschema_description:"The operation: add, subtract, multiply, or divide"`
}

// Weather
type WeatherInput struct {
	City string `json:"city" jsonschema_description:"The city to get weather for, e.g. 'Los Angeles'"`
}

type WeatherOutput struct {
	City        string  `json:"city"`
	TempC       float64 `json:"temp_celsius"`
	WindKph     float64 `json:"wind_kph"`
	Description string  `json:"description"`
}

// NFL Games
type NFLGamesInput struct {
	Placeholder string `json:"_" jsonschema_description:"Unused"`
}

type NFLGame struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// Moon Phase
type MoonPhaseInput struct {
	City string `json:"city" jsonschema_description:"City to reference (for context only — moon phase is the same globally at a given time)"`
}

type MoonPhaseOutput struct {
	City  string  `json:"city"`
	Phase string  `json:"phase"`
	Age   float64 `json:"age_days"` // days since new moon
}
