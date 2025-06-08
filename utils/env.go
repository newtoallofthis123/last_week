package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	GeminiApiKey string
}

// Reads the .env file and returns an Env struct.
func ReadEnv(localStore bool) Env {
	if !localStore {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(fmt.Errorf("error getting user home directory: %w", err))
		}

		godotenv.Load(home + "/.gcommit")
	} else {
		godotenv.Load(".env")
	}

	return Env{
		GeminiApiKey: getEnv("GOOGLE_GEMINI_KEY"),
	}
}

// Returns the value of the given env var name.
func getEnv(name string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("Env var %s not found", name))
	}
	return val
}
