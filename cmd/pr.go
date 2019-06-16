package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aflashyrhetoric/backlog-cli/utils"

	"net/url"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// PullRequest .. a PARTIAL struct for a PullRequest on Backlog
type PullRequest struct {
	Number      int               `json:"number"`
	Summary     string            `json:"summary"`
	Description string            `json:"description"`
	Base        string            `json:"base"`
	Branch      string            `json:"branch"`
	Issue       Issue             `json:"issue"`
	Status      pullRequestStatus `json:"status"`
}

type pullRequestStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var BaseBranch string

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

		// Get Issue Data
		endpoint := Endpoint(apiURL)
		responseData := utils.Get(endpoint)

		var currentIssue Issue
		json.Unmarshal(responseData, &currentIssue)
		// Convert integer -> string for use in later functions
		GlobalConfig.setIssue(currentIssue)

		// Create the form, request, and send the POST request
		// ---------------------------------------------------------
		p := GlobalConfig.ProjectKey
		r := GlobalConfig.RepositoryName
		apiURL = "/api/v2/projects/" + p + "/git/repositories/" + r + "/pullRequests"

		//apiURL = "test"
		endpoint = Endpoint(apiURL)

		existingPRs, err := checkForExistingPullRequests(endpoint)

		if len(existingPRs) > 0 {
			listPRs(existingPRs)
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("\n\nPull Requests for this issue already exist - would you still like to create one? (y\\n)")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)

			if text != "y" {
				return
			}
		}

		// Proceed with PR creation
		if CurrentBranch == "staging" || CurrentBranch == "dev" || CurrentBranch == "develop" || CurrentBranch == "beta" {
			fmt.Printf("CAUTION: You are on %s.\n", CurrentBranch)
			fmt.Printf("Creating PR: %s --> %s branch.\n", CurrentBranch, BaseBranch)
		} else if CurrentBranch == "0" {
			fmt.Println("Invalid branch. Try again.")
		} else {
			fmt.Printf("Creating PR: %s --> %s branch.\n", CurrentBranch, BaseBranch)
		}

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
		if GlobalConfig.CurrentIssue.ID != 0 {
			form.Add("issueId", strconv.Itoa(GlobalConfig.CurrentIssue.ID))
		}

		responseData, err = utils.Post(endpoint, form)
		ErrorPanic(err)

		var returnedPullRequest PullRequest
		json.Unmarshal(responseData, &returnedPullRequest)

		linkToPR := getPRLink(returnedPullRequest.Number)
		fmt.Printf("Link to PR: %s", linkToPR)
	},
}

func listPRs(PRList []PullRequest) {
	fmt.Printf("\n\n[Existing Pull Requests found]\n")
	for i, pr := range PRList {

		// If there are open PRs with a matching issue ID
		if pr.Status.ID == 1 && pr.Issue.ID == GlobalConfig.CurrentIssue.ID {
			fmt.Printf("\t%v: %s\n", i+1, getPRLink(pr.Number))
		}
	}

}

func getPRLink(n int) string {
	return fmt.Sprintf("%s/git/%s/%s/pullRequests/%d", GlobalConfig.BaseURL, GlobalConfig.ProjectKey, GlobalConfig.RepositoryName, n)
}

func checkForExistingPullRequests(endpoint string) ([]PullRequest, error) {
	responseData := utils.Get(endpoint)

	// List of Pull Requests that already exist and share the ID
	var existingPullRequests []PullRequest

	// List of returned pull requests
	var returnedPullRequests []PullRequest
	err := json.Unmarshal(responseData, &returnedPullRequests)

	ErrorCheck(err)

	for _, element := range returnedPullRequests {
		// fmt.Println(GlobalConfig.CurrentIssue)
		// fmt.Println(element)
		if element.Issue.ID == GlobalConfig.CurrentIssue.ID {
			existingPullRequests = append(existingPullRequests, element)
		}
	}

	return existingPullRequests, err
}

func init() {
	prCmd.Flags().StringVarP(&BaseBranch, "branch", "b", "master", "Designate a branch (other than master) to merge to.")
	RootCmd.AddCommand(prCmd)
}
