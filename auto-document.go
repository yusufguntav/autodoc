package autodocument

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func GenerateDocumentation() {

	var output string

	var rootCmd = &cobra.Command{Use: "autodoc"}

	var autodocCmd = &cobra.Command{
		Use:   "autodoc",
		Short: "Create documentation for your project",
		Run: func(cmd *cobra.Command, args []string) {
			err := godotenv.Load()
			if err != nil {
				fmt.Printf("error loading .env file: %v\n", err)
				return
			}

			branch, err := GetCurrentBranch()
			if err != nil {
				fmt.Println(err)
				return
			}

			commits, err := GetUnpushedCommits(branch)
			if err != nil {
				fmt.Println(err)
				return
			}

			result := SendMessageAI(commits)

			if output != "" {
				err := os.WriteFile(output, []byte(result), 0644)
				if err != nil {
					fmt.Printf("error writing to file: %v\n", err)
					return
				}
				fmt.Printf("Documentation written to %s\n", output)
			} else {
				fmt.Println(result)
			}

		},
	}

	autodocCmd.Flags().StringVarP(&output, "output", "o", "", "File path to write the documentation output")
	rootCmd.AddCommand(autodocCmd)
	rootCmd.Execute()
}
