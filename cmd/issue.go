package cmd

import (
	"encoding/json"
	"net/url"
	"regexp"

	"github.com/aflashyrhetoric/backlog-cli/utils"
	"github.com/spf13/cobra"
)

// Issue .. the configuration struct for backlog-cli
type Issue struct {
	ID          int    `json:"id"`
	IssueKey    string `json:"issueKey"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

var issueCmd = &cobra.Command{
	Use:     "issue",
	Aliases: []string{"issues", "i"},
	Short:   "Create issues",
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		endpoint := IssueListEndpoint()

		// 		templates := &promptui.SelectTemplates{
		// 			Label:    "{{ .Sender.Name }}",
		// 			Active:   "-> [{{ .Sender.Name | cyan }}] ",
		// 			Inactive: "   [{{ .Sender.Name | cyan }}] ",
		// 			Details: `
		// --- [{{ .Sender.Name | faint }}] ---
		// {{ .Content }}
		// ------------------------------------`,
		// 		}

		// notificationURL := fmt.Sprintf("%s/globalbar/notifications/redirect/%d", GlobalConfig.BaseURL, returnedNotifs[i].ID)
		// openBrowser(notificationURL)
		responseData, err := utils.Post(endpoint, url.Values{})
		ErrorCheck(err)
		var returnedNotifs []Notification
		json.Unmarshal(responseData, &returnedNotifs)
	},
}

func init() {
	// notifCmd.Flags().StringVarP(&BaseBranch, "branch", "b", "master", "Designate a branch (other than master) to merge to.")
	// RootCmd.AddCommand(issueCmd)
}

// Key .. Returns the issue's key (e.g. "BLG-123") as a string
func (i *Issue) Key() string {
	endpoint := IssueListEndpoint()
	responseData := utils.Get(endpoint)

	// A Response struct to map the Entire Response
	var returnedIssue Issue
	json.Unmarshal(responseData, &returnedIssue)

	return returnedIssue.IssueKey
}

// GetCurrentIssue .. Returns the current issue
func GetCurrentIssue() Issue {

	var issueID string

	// By default, get Issue ID from current branch name if possible
	cb := GlobalConfig.CurrentBranch
	reg := regexp.MustCompile("([a-zA-Z]+-[0-9]*)")
	if reg.Find([]byte(cb)) != nil {
		issueID = string(reg.Find([]byte(cb)))
	}

	endpoint := IssueEndpoint(issueID)
	responseData := utils.Get(endpoint)

	var currentIssue Issue
	json.Unmarshal(responseData, &currentIssue)

	return currentIssue
}

// GetProjectKey ... Returns the project key for the configuration (e.g "MARKETING")
func GetProjectKey() string {
	repo := GetCurrentRepo()

	originRemote, err := repo.Remote("origin")
	ErrorCheck(err)

	// Fetch references from 'origin'
	repoReferences := originRemote.String()

	// Capture repository name from reference
	reg := regexp.MustCompile(`\/([A-Z]*)\/(.*)\.git`)
	repositoryCapturedString := reg.FindSubmatch([]byte(repoReferences))
	projectKey := string(repositoryCapturedString[1])

	return projectKey
}
