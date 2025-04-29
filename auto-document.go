package autodocument

import (
	"fmt"

	"github.com/joho/godotenv"
)

const geminiAPIURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

// GenerateDocumentation is the main entry point for generating documentation
func GenerateDocumentation() (string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading .env file: %v\n", err)
	}

	branch, err := GetCurrentBranch()
	if err != nil {
		return "", fmt.Errorf("error fetching current branch: %w", err)
	}

	commits, err := GetCommits(branch)
	if err != nil {
		return "", fmt.Errorf("error fetching commits: %w", err)
	}

	return SendMessageAI(commits), nil
}
