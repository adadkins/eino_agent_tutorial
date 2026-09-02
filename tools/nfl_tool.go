package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NflGamesToday(ctx context.Context, input *NFLGamesInput) ([]NFLGame, error) {
	url := "https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch NFL scoreboard: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var espn struct {
		Events []struct {
			Name   string `json:"name"`
			Status struct {
				Type struct {
					Description string `json:"description"`
				} `json:"type"`
			} `json:"status"`
		} `json:"events"`
	}
	if err := json.Unmarshal(body, &espn); err != nil {
		return nil, fmt.Errorf("failed to parse NFL scoreboard response: %w", err)
	}

	games := make([]NFLGame, 0, len(espn.Events))
	for _, e := range espn.Events {
		games = append(games, NFLGame{Name: e.Name, Status: e.Status.Type.Description})
	}
	return games, nil
}
