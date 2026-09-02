package main

import (
	"context"
	"eino_agent_tutorial/tools"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
)

func buildAgent(ctx context.Context) (*react.Agent, error) {
	llmURL := os.Getenv("LLM_URL")
	llmModel := os.Getenv("LLM_MODEL")
	llmAPIKey := os.Getenv("LLM_API_KEY")
	// make a new chat model pointed to my home domain
	cm, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: llmURL,
		Model:   llmModel,
		APIKey:  llmAPIKey,
	})
	if err != nil {
		return nil, err
	}
	// register our tools so the model can call them
	calcTool, err := utils.InferTool("calculator", "Performs basic arithmetic operations: add, subract, multiply, divide", tools.Calculator)
	if err != nil {
		return nil, err
	}
	weatherTool, err := utils.InferTool("weather", "Provides weather information for a given location", tools.Weather)
	if err != nil {
		return nil, err
	}

	nflTool, err := utils.InferTool("nfl", "Get todays' NFL games and their status", tools.NflGamesToday)
	if err != nil {
		return nil, err
	}
	moonTool, err := utils.InferTool("moon", "Gets the current moon phase, optionally for a specific city. Defaults to Los Angeles", tools.MoonPhase)
	if err != nil {
		return nil, err
	}

	return react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: cm,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: []tool.BaseTool{calcTool, weatherTool, nflTool, moonTool},
		},
	})
}
