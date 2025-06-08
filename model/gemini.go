package model

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/newtoallofthis123/last_week/utils"
	"google.golang.org/api/option"
)

type GeminiModel struct {
	client *genai.Client
	ctx    context.Context
}

func defaultSystemPrompt() string {
	return `
	You are a helpful assistant that can answer questions and help with tasks.
	`
}

func NewGeminiModel(ctx context.Context, env *utils.Env) (*GeminiModel, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(env.GeminiApiKey))
	if err != nil {
		return nil, err
	}
	return &GeminiModel{client: client, ctx: ctx}, nil
}

func (m *GeminiModel) GetModel(systemPrompt string) *genai.GenerativeModel {
	finalSystemPrompt := ""
	if systemPrompt != "" {
		finalSystemPrompt = systemPrompt
	} else {
		finalSystemPrompt = defaultSystemPrompt()
	}

	model := m.client.GenerativeModel("gemini-2.0-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(finalSystemPrompt)},
	}
	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)

	return model
}

func (m *GeminiModel) GenerateResponse(prompt string, systemPrompt string) (string, error) {
	model := m.GetModel(systemPrompt)
	modelOutput, err := model.GenerateContent(m.ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content from AI model: %w", err)
	}
	finalOutput := ""

	for _, cand := range modelOutput.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				finalOutput += fmt.Sprintf("%s", part)
			}
		}
	}

	return finalOutput, nil
}
