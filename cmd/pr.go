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
	Number      int    `json:"number"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Base        string `json:"base"`
	Branch      string `json:"branch"`
}

var branchName string
var currentIssue Issue

// Gets
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates a Backlog Pull Request for the current branch -> master",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		// By default, get Issue ID from current branch name if possible
		// var formContentType := "Content-Type:application/x-www-form-urlencoded"
		CurrentBranch := CurrentBranch()
		reg := regexp.MustCompile("([a-zA-Z]+-[0-9]*)")
		issueID := string(reg.Find([]byte(CurrentBranch)))
		fmt.Println(issueID)

		apiURL := "/api/v2/issues/" + string(issueID)

		if CurrentBranch == "staging" || CurrentBranch == "dev" || CurrentBranch == "develop" || CurrentBranch == "beta" {
			fmt.Printf("You're currently on the %v branch. Please switch to an issue branch and try again.", issueID)
		} else if CurrentBranch == "0" {
			fmt.Println("Invalid branch. Try again.")
		} else {
			fmt.Printf("Creating PR for %v branch.", CurrentBranch)
		}

		endpoint := Endpoint(apiURL)
		// fmt.Println("Endpoint is:")
		// fmt.Println(endpoint)
		responseData := utils.Get(endpoint)
		json.Unmarshal(responseData, &currentIssue)
		// Convert integer -> string for use in later functions
		issueID = strconv.Itoa(currentIssue.ID)

		// Create the form, request, and send the POST request
		// ---------------------------------------------------------
		p := ProjectKey()
		r := RepositoryName()
		apiURL = "/api/v2/projects/" + p + "/git/repositories/" + r + "/pullRequests"

		//apiURL = "test"
		endpoint = Endpoint(apiURL)

		// Build out Form
		form := url.Values{}
		form.Add("summary", "Test summary")
		form.Add("description", "Test description")
		// Branch to merge to
		form.Add("base", branchName)
		// Branch of branch we are merging
		form.Add("branch", CurrentBranch)
		form.Add("issueId", issueID)
		form.Add("assigneeId", strconv.Itoa(GlobalConfig.User.ID))

		responseData = utils.Post(endpoint, form)

		var returnedPullRequest pullRequest
		json.Unmarshal(responseData, &returnedPullRequest)

		fmt.Println(returnedPullRequest)
		currentPullRequestID := strconv.Itoa(returnedPullRequest.Number)

		linkToPR := fmt.Sprintf("%s/git/%s/%s/pullRequests/%s", GlobalConfig.BaseURL, GlobalConfig.ProjectKey, GlobalConfig.RepositoryName, currentPullRequestID)
		fmt.Printf("Link to PR: %s", linkToPR)
	},
}

func init() {
	prCmd.Flags().StringVarP(&branchName, "branch", "b", "master", "Designate a branch (other than master) to merge to.")
	RootCmd.AddCommand(prCmd)
}
