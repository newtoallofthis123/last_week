package model

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

type OllamaModel struct {
	client *api.Client
	ctx    context.Context
}

func NewOllamaModel(ctx context.Context) (*OllamaModel, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, err
	}
	return &OllamaModel{client, ctx}, nil
}

func (m *OllamaModel) GenerateResponse(prompt string, systemPrompt string) (string, error) {
	if systemPrompt == "" {
		systemPrompt = defaultSystemPrompt()
	}
	stream := false
	request := &api.GenerateRequest{
		Model:  "gemma3:latest",
		System: systemPrompt,
		Stream: &stream,
	}

	request.Prompt = prompt
	var response string
	respFunc := func(resp api.GenerateResponse) error {
		response = resp.Response
		return nil
	}
	err := m.client.Generate(m.ctx, request, respFunc)
	if err != nil {
		return "", fmt.Errorf("failed to generate content from AI model: %w", err)
	}

	return response, nil
}
