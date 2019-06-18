package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aflashyrhetoric/backlog-cli/utils"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

// Notification .. a PARTIAL struct for a Notification
type Notification struct {
	ID                 int                            `json:"id"`
	Sender             notificationSender             `json:"sender"`
	Comment            notificationComment            `json:"comment"`
	PullRequestComment notificationPullRequestComment `json:"pullRequestComment"`
	Reason             int                            `json:"reason"`
	Content            string
}

// NotificationSender..
type notificationSender struct {
	Name string `json:"name"`
}
type notificationComment struct {
	Content string `json:"content"`
}
type notificationPullRequestComment struct {
	Content string `json:"content"`
}

var notifCmd = &cobra.Command{
	Use:   "n",
	Short: "Read and open your notifications",
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		apiURL := "/api/v2/notifications"

		endpoint := Endpoint(apiURL)

		params := map[string]int{
			"count": 15,
		}
		// Add issueID if it exists
		responseData := utils.GetParams(endpoint, params)

		var returnedNotifs []Notification
		json.Unmarshal(responseData, &returnedNotifs)

		var transformedNotifs []Notification

		for _, n := range returnedNotifs {
			n.truncateName()
			transformedNotifs = append(transformedNotifs, n)
		}

		templates := &promptui.SelectTemplates{
			Label:    "{{ .Sender.Name }}",
			Active:   "-> [{{ .Sender.Name | cyan }}] ",
			Inactive: "   [{{ .Sender.Name | cyan }}] ",
			Details: `
--- [{{ .Sender.Name | faint }}] ---
{{ .Content }}
------------------------------------`,
		}

		prompt := promptui.Select{
			Label:     "Select notification",
			Items:     transformedNotifs,
			Templates: templates,
			Size:      6,
		}

		i, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		// fmt.Printf("You choose number %d: %s\n", i+1, returnedNotifs[i].Sender.Name)
		notificationURL := fmt.Sprintf("%s/globalbar/notifications/redirect/%d", GlobalConfig.BaseURL, returnedNotifs[i].ID)
		openBrowser(notificationURL)
	},
}

func (n *Notification) truncateName() {
	var text string
	maxCharacterCount := 60
	switch n.Reason {
	case 1:
		text = n.Sender.Name + " changed the issue's assignee to you."
		break
	case 2:
		text = n.Comment.Content
		break
	case 3:
		text = n.Sender.Name + " added an issue."
		break
	case 11:
		text = n.PullRequestComment.Content
		break
	case 13:
		text = n.PullRequestComment.Content
		break
	default:
		text = "Issue type unknown"
		break
	}
	if len(text) > maxCharacterCount {
		n.Content = text[:maxCharacterCount]
	} else {
		n.Content = text
	}
}

func init() {
	// notifCmd.Flags().StringVarP(&BaseBranch, "branch", "b", "master", "Designate a branch (other than master) to merge to.")
	RootCmd.AddCommand(notifCmd)
}
