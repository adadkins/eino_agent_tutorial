package tools

import (
	"context"
	"testing"
)

func TestNflGamesToday(t *testing.T) {
	games, err := NflGamesToday(context.Background(), &NFLGamesInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// games can legitimately be empty (no games today) — just check it doesn't error
	_ = games
}
