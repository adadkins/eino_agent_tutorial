package main

import (
	"context"
	"testing"
)

func TestBuildAgent_Success(t *testing.T) {
	agent, err := buildAgent(context.Background())
	if err != nil {
		t.Fatalf("expected buildAgent to succeed, got error: %v", err)
	}
	if agent == nil {
		t.Fatal("expected a non-nil agent")
	}
}
