package model

type Model interface {
	GenerateResponse(prompt string, systemPrompt string) (string, error)
}
