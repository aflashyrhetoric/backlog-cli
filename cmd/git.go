package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	git "gopkg.in/src-d/go-git.v4"
)

var gitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long:  `to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		r := Repo()
		logs, err := r.Log(&git.LogOptions{})
		ErrorPanic(err)

		latestCommit, err := logs.Next()
		ErrorPanic(err)

		logs.Close()
		latestCommitID := latestCommit.ID()

		templateURL := fmt.Sprintf("%s/git/%s/%s/commit/%v", viper.GetString("BASEURL"), ProjectKey(), RepoName(), latestCommitID)
		// templateURL :=

		// Fetch
		fmt.Printf("Your latest commit: %s", templateURL)
	},
}

func init() {
	RootCmd.AddCommand(gitCmd)
}
