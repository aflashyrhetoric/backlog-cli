package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
)

var gitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Get a link to the latest commit",
	Long:  `to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		// Open repository and fetch logs
		r := Repository()
		logs, err := r.Log(&git.LogOptions{})
		ErrorPanic(err)

		// Fetch latest commit
		latestCommit, err := logs.Next()
		ErrorPanic(err)
		logs.Close()

		// Fetch ID from latest commit
		latestCommitID := latestCommit.ID()

		// Assemble the URL
		commitURL := fmt.Sprintf("%s/git/%s/%s/commit/%v", GlobalConfig.BaseURL, ProjectKey(), RepositoryName(), latestCommitID)

		// Fetch
		fmt.Printf("Your latest commit: %s", commitURL)
	},
}

func init() {
	RootCmd.AddCommand(gitCmd)
}
