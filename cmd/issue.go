package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/aflashyrhetoric/backlog-cli/utils"
	"github.com/spf13/cobra"
)

// Issue .. the configuration struct for backlog-cli
type Issue struct {
	ID          int    `json:"id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	IssueKey    string `json:"issueKey"`
	AssigneeID  User   `json:"assigneeId"`
}

var issueCmd = &cobra.Command{
	Use: "issue",
	// Hidden:  true,
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

var createIssueCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create issues",
	Run: func(cmd *cobra.Command, args []string) {

		// projectId = 2
		// issueTypeId= 2
		// priorityId= 3
		fmt.Println("BLG: Creating new issue.")
		// Create issue Form
		form := url.Values{}

		assignee, err := PromptInput("Type the name of an assignee.")
		assignee = strings.ToLower(assignee)
		ErrorCheck(err)

		if assignee == "me" {
			assignee = strconv.Itoa(GlobalConfig.User.ID)
			fmt.Println(assignee)
		} else {
			matchedUsers := searchUsers(assignee, GetUserList())
			// If match not found
			if matchedUsers == nil {
				fmt.Println("Match not found, please try again.")
				return
			}
			if len(matchedUsers) == 1 {
				assignee = string(matchedUsers[0].ID)
			}

			if len(matchedUsers) > 1 {
				matchedUser := AssigneeSelect(matchedUsers)
				assignee = strconv.Itoa(matchedUser.ID)
			}
		}

		// Addknown values + defaults
		form.Add("assigneeId", assignee)
		form.Add("issueTypeId", "2")
		form.Add("priorityId", "3")
		form.Add("projectId", "1073846676")
		form.Add("projectKey", "EXPERIMENT")

		summary, err := AskFormField("Summary")
		ErrorCheck(err)
		form.Add("summary", summary)

		description, err := AskFormField("Description")
		ErrorCheck(err)
		form.Add("description", description)

		endpoint := IssueListEndpoint()
		responseData, err := utils.Post(endpoint, form)
		ErrorPanic(err)

		var returnedIssue Issue
		json.Unmarshal(responseData, &returnedIssue)

		linkToIssue := LinkToIssuePage(returnedIssue.Key())
		fmt.Printf("Link to Issue: %s", linkToIssue)
	},
}

func searchUsers(userToFind string, users []User) []User {

	var matches []User

	for _, user := range users {
		userName := user.Name
		userUsername := user.Username

		//
		if strings.Contains(userName, userToFind) || strings.Contains(userUsername, userToFind) {
			matches = append(matches, user)
		}
	}
	if len(matches) > 0 {
		return matches
	}

	return nil
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

func init() {
	// notifCmd.Flags().StringVarP(&BaseBranch, "branch", "b", "master", "Designate a branch (other than master) to merge to.")
	issueCmd.AddCommand(createIssueCmd)
	RootCmd.AddCommand(issueCmd)
}
