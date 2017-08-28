package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"net/url"
	"strconv"
)

// Gets
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates a Backlog Pull Request for the current branch",
	Long:  `abc`,
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------
		i := "CACOO-11747"
		apiUrl := "issues/" + i
		endpoint := Endpoint(apiUrl)
		type Issue struct {
			Id int `json:"id"`
		}
		responseData := get(endpoint)
		var currentIssue Issue
		json.Unmarshal(responseData, &currentIssue)
		issueId := strconv.Itoa(currentIssue.Id)

		fmt.Println(issueId)

		// ---------------------------------------------------------
		p := ProjectKey()
		r := Repo()

		apiUrl = "projects/" + p + "/git/repositories/" + r + "/pullRequests"
		endpoint = Endpoint(apiUrl)

		// Build out Form
		form := url.Values{}
		form.Add("summary", "Test summary")
		form.Add("description", "Test description")
		form.Add("base", "master")
		form.Add("branch", "test")
		form.Add("issueId", string(issueId))

		responseData = post(endpoint, form)

		// A Response struct to map the Entire Response
		type PullRequest struct {
			Summary     string `json:"summary"`
			Description string `json:"description"`
			Base        string `json:"base"`
			Branch      string `json:"branch"`
			Id          int    `json:"id"`
		}

		var returnedPullRequestCount PullRequest

		json.Unmarshal(responseData, &returnedPullRequestCount)

		printResponse(responseData)

		// Failure
		fmt.Println(returnedPullRequestCount.Summary)

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
	//fmt.Printf("%s\n", refer.Name())
	return nil
}
