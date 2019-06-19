package cmd

import (
	"encoding/json"
	"net/url"
	"strconv"

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
	Aliases: []string{"i"},
	Short:   "Create issues",
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		apiURL := "/api/v2/issues"

		endpoint := Endpoint(apiURL)

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
	RootCmd.AddCommand(issueCmd)
}

// Key .. Returns the issue's key (e.g. "BLG-123") as a string
func (i *Issue) Key() string {
	apiURL := "/api/v2/issues/" + strconv.Itoa(i.ID)
	endpoint := Endpoint(apiURL)
	responseData := utils.Get(endpoint)

	// A Response struct to map the Entire Response
	var returnedIssue Issue
	json.Unmarshal(responseData, &returnedIssue)

	return returnedIssue.IssueKey
}
