package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aflashyrhetoric/backlog-cli/utils"

	"net/url"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// PullRequest .. a PARTIAL struct for a PullRequest on Backlog
type PullRequest struct {
	Number      int    `json:"number"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Base        string `json:"base"`
	Branch      string `json:"branch"`
}

var BaseBranch string
var currentIssue Issue

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates a Backlog Pull Request for the current branch to (master) or some other branch",
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		// By default, get Issue ID from current branch name if possible
		// var formContentType := "Content-Type:application/x-www-form-urlencoded"
		CurrentBranch := CurrentBranch()
		reg := regexp.MustCompile("([a-zA-Z]+-[0-9]*)")

		var issueID string

		if reg.Find([]byte(CurrentBranch)) != nil {
			issueID = string(reg.Find([]byte(CurrentBranch)))
		}

		apiURL := "/api/v2/issues/" + string(issueID)

		if CurrentBranch == "staging" || CurrentBranch == "dev" || CurrentBranch == "develop" || CurrentBranch == "beta" {
			fmt.Printf("CAUTION: You are on %s.\n", CurrentBranch)
			fmt.Printf("Creating PR: %s --> %s branch.\n", CurrentBranch, BaseBranch)
		} else if CurrentBranch == "0" {
			fmt.Println("Invalid branch. Try again.")
		} else {
			fmt.Printf("Creating PR: %s --> %s branch.\n", CurrentBranch, BaseBranch)
		}

		endpoint := Endpoint(apiURL)
		responseData := utils.Get(endpoint)
		json.Unmarshal(responseData, &currentIssue)
		// Convert integer -> string for use in later functions
		issueID = strconv.Itoa(currentIssue.ID)

		// Create the form, request, and send the POST request
		// ---------------------------------------------------------
		p := GlobalConfig.ProjectKey
		r := GlobalConfig.RepositoryName
		apiURL = "/api/v2/projects/" + p + "/git/repositories/" + r + "/pullRequests"

		//apiURL = "test"
		endpoint = Endpoint(apiURL)

		// Build out Form
		form := url.Values{}
		form.Add("summary", fmt.Sprintf("Test summary: %v", time.Now()))
		form.Add("description", "Test description")
		// Branch to merge to
		form.Add("base", BaseBranch)
		// Branch of branch we are merging
		form.Add("branch", CurrentBranch)
		form.Add("assigneeId", strconv.Itoa(GlobalConfig.User.ID))

		// Add issueID if it exists
		if issueID != "0" {
			form.Add("issueId", issueID)
		}

		responseData, err := utils.Post(endpoint, form)
		ErrorPanic(err)

		var returnedPullRequest PullRequest
		json.Unmarshal(responseData, &returnedPullRequest)

		currentPullRequestID := strconv.Itoa(returnedPullRequest.Number)

		linkToPR := fmt.Sprintf("%s/git/%s/%s/pullRequests/%s", GlobalConfig.BaseURL, GlobalConfig.ProjectKey, GlobalConfig.RepositoryName, currentPullRequestID)
		fmt.Printf("Link to PR: %s", linkToPR)
	},
}

func init() {
	prCmd.Flags().StringVarP(&BaseBranch, "branch", "b", "master", "Designate a branch (other than master) to merge to.")
	RootCmd.AddCommand(prCmd)
}
