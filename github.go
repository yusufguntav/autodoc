package autodocument

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// GetCurrentBranch fetches the active branch name
func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// GetUnpushedCommits fetches unpushed commit hashes and messages for the active branch
func GetUnpushedCommits(branch string) (string, error) {
	cmd := exec.Command("git", "log", fmt.Sprintf("origin/%s..HEAD", branch), "--pretty=format:%H")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	commitHashes := strings.Split(strings.TrimSpace(string(output)), "\n")
	commitChanges := make(map[string]string)

	for _, hash := range commitHashes {
		cmd := exec.Command("git", "show", hash, "--pretty=format:", "--unified=0")
		diffOutput, err := cmd.Output()
		if err != nil {
			return "", err
		}
		commitChanges[hash] = string(diffOutput)
	}

	jsonOutput, err := json.MarshalIndent(commitChanges, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonOutput), nil
}
