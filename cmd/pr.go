package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aflashyrhetoric/backlog-cli/utils"

	"net/url"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

type pullRequest struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Base        string `json:"base"`
	Branch      string `json:"branch"`
	ID          int    `json:"id"`
}

// Gets
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates a Backlog Pull Request for the current branch -> master",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		// By default, get Issue ID from current branch name if possible
		CurrentBranch := CurrentBranch()
		reg := regexp.MustCompile("([a-zA-Z]+-[0-9]*)")
		issueID := string(reg.Find([]byte(CurrentBranch)))
		fmt.Println(issueID)

		apiURL := "/api/v2/issues/" + string(issueID)

		if CurrentBranch == "staging" || CurrentBranch == "dev" || CurrentBranch == "beta" {
			fmt.Printf("You're currently on the %v branch. Please switch to an issue branch and try again.", issueID)
		} else if CurrentBranch == "0" {
			fmt.Println("Invalid branch. Try again.")
		} else {
			fmt.Printf("Creating PR for %v branch.", CurrentBranch)
		}

		endpoint := Endpoint(apiURL)
		fmt.Println("Endpoint is:")
		fmt.Println(endpoint)
		type Issue struct {
			ID int `json:"id"`
		}
		responseData := utils.Get(endpoint)
		var currentIssue Issue
		json.Unmarshal(responseData, &currentIssue)
		// Convert integer -> string for use in later functions
		issueID = strconv.Itoa(currentIssue.ID)

		// Create the form, request, and send the POST request
		// ---------------------------------------------------------
		p := ProjectKey()
		r := RepoName()
		apiURL = "/api/v2/projects/" + p + "/git/repositories/" + r + "/pullRequests"

		//apiURL = "test"
		endpoint = Endpoint(apiURL)

		// Build out Form
		form := url.Values{}
		form.Add("summary", "Test summary")
		form.Add("description", "Test description")
		form.Add("base", "master")
		form.Add("branch", CurrentBranch)
		form.Add("issueId", issueID)

		responseData = utils.Post(endpoint, form)

		var returnedPullRequestCount pullRequest
		json.Unmarshal(responseData, &returnedPullRequestCount)
		PrintResponse(responseData)

		fmt.Println(returnedPullRequestCount.Summary)
	},
}

func init() {
	RootCmd.AddCommand(prCmd)
}
