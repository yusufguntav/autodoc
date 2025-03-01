package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// getCommits fetches unpushed commit hashes and messages
func getCommits() ([]string, error) {
	cmd := exec.Command("git", "log", "origin/master..HEAD", "--pretty=format:%H %s") // Sadece pushlanmamış commitleri al
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	return lines, nil
}

// getChangedFiles fetches changed files for a given commit
func getChangedFiles(commitHash string) ([]string, error) {
	cmd := exec.Command("git", "show", "--name-only", "--pretty=format:", commitHash)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	files := strings.Split(strings.TrimSpace(string(output)), "\n")
	return files, nil
}

// classifyFiles separates frontend and general changes
func classifyFiles(files []string) (frontend []string, general []string) {
	frontendExts := []string{".js", ".jsx", ".ts", ".tsx", ".css", ".scss", ".html"}
	for _, file := range files {
		isFrontend := false
		for _, ext := range frontendExts {
			if strings.HasSuffix(file, ext) {
				frontend = append(frontend, file)
				isFrontend = true
				break
			}
		}
		if !isFrontend {
			general = append(general, file)
		}
	}
	return frontend, general
}

func main() {
	commits, err := getCommits()
	if err != nil {
		fmt.Println("Error fetching commits:", err)
		return
	}

	for _, commit := range commits {
		parts := strings.SplitN(commit, " ", 2)
		if len(parts) < 2 {
			continue
		}
		hash, message := parts[0], parts[1]
		files, err := getChangedFiles(hash)
		if err != nil {
			fmt.Println("Error fetching changed files for commit", hash, err)
			continue
		}

		frontend, general := classifyFiles(files)
		fmt.Printf("Commit: %s\nMessage: %s\nFrontend Changes: %v\nGeneral Changes: %v\n\n", hash, message, frontend, general)
	}
}
