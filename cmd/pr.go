package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Gets
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates a Backlog Pull Request for the current branch",
	Long:  `abc`,
	Run: func(cmd *cobra.Command, args []string) {

		p := ProjectKey()
		r := Repo()

		// FIXME: Temporary way to build up api endpoint url
		apiUrl := "projects/" + p + "/git/repositories/" + r + "/pullRequests/count"
		endpoint := Endpoint(apiUrl)

		// Fetch
		responseData := get(endpoint)

		// TODO: Add struct to map out JSON response for PR
		// TODO: Unmarshal and populate JSON properly

		// A Response struct to map the Entire Response
		type PullRequest struct {
			Count int8 `json:"count"`
		}

		var returnedPullRequestCount PullRequest

		json.Unmarshal(responseData, &returnedPullRequestCount)
		fmt.Println(returnedPullRequestCount.Count)

		repo, err := git.PlainOpen(path)

		if err != nil {
			fmt.Printf("#%v", err)
			panic(err)
		}
		branches, err := repo.Branches()
		if err != nil {
			panic(err)
		}

		err = branches.ForEach(reference)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(prCmd)
}

func reference(refer *plumbing.Reference) error {
	//fmt.Printf("%#v\n", refer)
	//fmt.Printf("%s\n", refer.Count)
	return nil
}
