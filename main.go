package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const geminiAPIURL = "https://generativelanguage.googleapis.com/v1/models/gemini-pro:generateText"

type GeminiRequest struct {
	Prompt string `json:"prompt"`
}

type GeminiResponse struct {
	Candidates []struct {
		Output string `json:"output"`
	} `json:"candidates"`
}

// getCurrentBranch fetches the active branch name
func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// getCommits fetches unpushed commit hashes and messages for the active branch
func getCommits(branch string) ([]string, error) {
	cmd := exec.Command("git", "log", fmt.Sprintf("origin/%s..HEAD", branch), "--pretty=format:%H %s") // Sadece pushlanmamış commitleri al
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	return lines, nil
}

// summarizeCommits uses Gemini AI to summarize commit messages
func summarizeCommits(messages []string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY environment variable is not set")
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"contents": []map[string]string{
			{"parts": strings.Join(messages, "\n")},
		},
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s?key=%s", geminiAPIURL, apiKey), bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if len(response.Candidates) > 0 {
		return response.Candidates[0].Output, nil
	}

	return "No summary generated", nil
}

func main() {
	branch, err := getCurrentBranch()
	if err != nil {
		fmt.Println("Error fetching current branch:", err)
		return
	}

	commits, err := getCommits(branch)
	if err != nil {
		fmt.Println("Error fetching commits:", err)
		return
	}

	var commitMessages []string
	for _, commit := range commits {
		parts := strings.SplitN(commit, " ", 2)
		if len(parts) < 2 {
			continue
		}
		hash, message := parts[0], parts[1]
		commitMessages = append(commitMessages, message)
		fmt.Printf("Commit: %s\nMessage: %s\n\n", hash, message)
	}

	if len(commitMessages) > 0 {
		summary, err := summarizeCommits(commitMessages)
		if err != nil {
			fmt.Println("Error summarizing commits:", err)
		} else {
			fmt.Println("\nSummary:", summary)
		}
	} else {
		fmt.Println("No new commits to summarize.")
	}
}
