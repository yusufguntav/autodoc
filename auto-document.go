package autodocument

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// GenerateDocumentation is the main entry point for generating documentation
func GenerateDocumentation() (string, error) {

	var rootCmd = &cobra.Command{Use: "autodoc"}

	var autodocCmd = &cobra.Command{
		Use:   "autodoc",
		Short: "Otomatik dokümantasyon üret",
		Run: func(cmd *cobra.Command, args []string) {
			err := godotenv.Load()
			if err != nil {
				fmt.Printf("error loading .env file: %v\n", err)
			}

			branch, err := GetCurrentBranch()
			if err != nil {
				fmt.Println(err)
			}

			commits, err := GetCommits(branch)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(SendMessageAI(commits))
		},
	}

	rootCmd.AddCommand(autodocCmd)
	rootCmd.Execute()
	return "", nil
}
