package main

import (
	"context"
	"eino_agent_tutorial/tools"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	einoagent "github.com/cloudwego/eino/flow/agent"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	callbackHelper "github.com/cloudwego/eino/utils/callbacks"
	"github.com/joho/godotenv"
)

const systemPrompt = `You are a narrow-purpose assistant. You may ONLY help with exactly these things:
1. Basic arithmetic calculations (use the calculator tool)
2. Current weather for a city (use the weather tool)
3. Today's NFL games (use the nfl_games_today tool)
4. The current moon phase for a city (use the moon_phase tool)
 
For ANY of these four topics, you MUST use the matching tool — never answer from your own knowledge.
For ANY other topic (general knowledge, opinions, coding help, anything unrelated to the four tools above),
you MUST refuse and respond with exactly: "I can only help with calculations, weather, NFL games, and moon phase questions."
Do not explain further, do not apologize at length, just give that exact refusal line.`

const refusalMessage = `I can only help with calculations, weather, NFL games, and moon phase questions.`

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using environment variables")
	}

	ctx := context.Background()

	agent, err := buildAgent(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to build agent:", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/chat", chatHandler(agent))

	port := "8090"
	log.Printf("agent API listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func chatHandler(agent *react.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req tools.ChatRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		if req.Prompt == "" {
			http.Error(w, "prompt is required", http.StatusBadRequest)
			return
		}
		// plug into the onstart hook to check if a tool was fired
		toolFired := false
		handler := callbackHelper.NewHandlerHelper().
			Tool(&callbackHelper.ToolCallbackHandler{
				OnStart: func(ctx context.Context, info *callbacks.RunInfo, input *tool.CallbackInput) context.Context {
					toolFired = true
					return ctx
				},
			}).Handler()

		messages := []*schema.Message{
			schema.SystemMessage(systemPrompt),
			schema.UserMessage(req.Prompt),
		}

		msg, err := agent.Generate(r.Context(), messages,
			einoagent.WithComposeOptions(compose.WithCallbacks(handler)),
		)
		if err != nil {
			log.Printf("agent error: %v", err)
			http.Error(w, "agent failed to generate response", http.StatusInternalServerError)
			return
		}

		answer := msg.Content

		//super basic check to make sure a tool was used
		if !toolFired {
			answer = refusalMessage
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tools.ChatResponse{Answer: answer, ToolUsed: toolFired})
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
