package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
)

var openFlag bool

// FIXME: 'commit' should be a child of `git`, or in its own file, etc
var gitCmd = &cobra.Command{
	Use:     "latest",
	Aliases: []string{"lc"},
	Short:   "Get a link to the latest commit",

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
		if openFlag {
			openBrowser(commitURL)
		}
	},
}

func init() {
	gitCmd.Flags().BoolVarP(&openFlag, "open", "o", false, "Include to open in a browser window")
	RootCmd.AddCommand(gitCmd)
}
