package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aflashyrhetoric/backlog-cli/utils"
	e "github.com/kyokomi/emoji"
	a "github.com/logrusorgru/aurora"

	"net/url"
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

// The branch that we wish to merge to
var BaseBranch string

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Create pull requests or open them (if there is only one associated pull request)",
}

var prCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Example: "pr create | pr create -b develop",
	Short:   "Creates a Backlog Pull Request for the current branch to (master) or some other branch",
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		cb := GlobalConfig.CurrentBranch

		sb := NewStringBuilder()
		endpoint := sb.PREndpoint()

		existingPRs, err := checkForExistingPullRequests(endpoint)

		// If there are PRs, warn the user
		if len(existingPRs) > 0 {
			listPRs("hand", "Existing Pull Requests found", existingPRs)
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("\n\nPull Requests for this issue already exist - would you still like to create one? (y\\n) ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if text != "y" {
				return
			}
		}

		// Proceed with PR creation
		if cb == "staging" || cb == "dev" || cb == "develop" || cb == "beta" {
			e.Printf(":rotating_light: CAUTION: You are on %s.\n", cb)
			fmt.Printf("Creating PR: %s --> %s branch.\n", cb, BaseBranch)
		} else if cb == "0" {
			fmt.Println("Invalid branch. Try again.")
		} else {
			e.Printf("Creating PR: %s --> %s branch. :zap: \n", cb, BaseBranch)
		}

		// Create the form, request, and send the POST request
		// ---------------------------------------------------------
		form := url.Values{}
		form.Add("summary", Truncate(GlobalConfig.CurrentIssue.Summary, 45))
		form.Add("description", GlobalConfig.CurrentIssue.Description)
		// Branch to merge to
		form.Add("base", BaseBranch)
		// Branch of branch we are merging
		form.Add("branch", cb)
		form.Add("assigneeId", strconv.Itoa(GlobalConfig.User.ID))

		// Add issueID if it exists
		if GlobalConfig.CurrentIssue.ID != 0 {
			form.Add("issueId", strconv.Itoa(GlobalConfig.CurrentIssue.ID))
		}

		responseData, err := utils.Post(endpoint, form)
		ErrorPanic(err)

		var returnedPullRequest PullRequest
		json.Unmarshal(responseData, &returnedPullRequest)

		linkToPR := getPRLink(returnedPullRequest.Number)
		fmt.Printf("Link to PR: %s", linkToPR)
	},
}
var prOpenCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o"},
	Short:   "If there's a single PR, open it in a browser.",
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		sb := NewStringBuilder()
		endpoint := sb.PREndpoint()

		existingPRs, err := checkForExistingPullRequests(endpoint)
		ErrorCheck(err)

		if len(existingPRs) == 1 {
			fmt.Println("Opening associated PR")
			openBrowser(getPRLink(existingPRs[0].Number))
		}
	},
}

var prListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List links to all associated pull requests",
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		sb := NewStringBuilder()
		endpoint := sb.PREndpoint()

		existingPRs, err := checkForExistingPullRequests(endpoint)
		ErrorCheck(err)

		if len(existingPRs) == 0 {
			fmt.Println("There are no pull requests associated to this issue.")
		}
		if len(existingPRs) > 0 {
			listPRs("sparkles", fmt.Sprintf("Pull Requests for %s", GlobalConfig.CurrentIssue.IssueKey), existingPRs)
		}

	},
}

func listPRs(emoji string, message string, PRList []PullRequest) {
	e.Print("\n\n:" + emoji + ":")
	fmt.Println(a.White("[" + message + "]").BgBrightBlack())
	count := 1
	for _, pr := range PRList {

		// If there are open PRs with a matching issue ID
		if pr.Status.ID == 1 && pr.Issue.ID == GlobalConfig.CurrentIssue.ID {
			fmt.Printf("%v: %s\n", count, getPRLink(pr.Number))
			fmt.Printf("   %s %s %s\n", a.Cyan(pr.Branch), a.Bold("-->"), a.Cyan(pr.Base))
			count++
		}
	}
}

func getPRLink(n int) string {
	return fmt.Sprintf("%s/git/%s/%s/pullRequests/%d", GlobalConfig.BaseURL, GlobalConfig.ProjectKey, GlobalConfig.RepositoryName, n)
}

func checkForExistingPullRequests(endpoint string) ([]PullRequest, error) {

	// params for pull requests
	params := map[string]int{
		"statusId[]": 1,
	}

	responseData := utils.GetParams(endpoint, params)

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
	prCreateCmd.Flags().StringVarP(&BaseBranch, "branch", "b", "master", "Designate a branch (other than master) to merge to.")
	prCmd.AddCommand(prCreateCmd)
	prCmd.AddCommand(prOpenCmd)
	prCmd.AddCommand(prListCmd)
	RootCmd.AddCommand(prCmd)
}
