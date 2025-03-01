package main

import (
	"fmt"
	"os/exec"
	"strings"
)

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
