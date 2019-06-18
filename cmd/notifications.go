package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aflashyrhetoric/backlog-cli/utils"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

// Notification .. a PARTIAL struct for a Notification
type Notification struct {
	ID          int                     `json:"id"`
	Sender      notificationSender      `json:"sender"`
	Comment     notificationComment     `json:"comment"`
	PullRequest notificationPullRequest `json:"pullRequestComment"`
	content     string
}

// NotificationSender..
type notificationSender struct {
	Name string `json:"name"`
}
type notificationComment struct {
	Content string `json:"content"`
}
type notificationPullRequest struct {
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

			if len(n.PullRequest.Content) > 0 {
				n.content = n.PullRequest.Content[:5]
			}
			if len(n.Comment.Content) > 0 {
				n.content = n.Comment.Content[:5]
			}

			fmt.Print(n)
			transformedNotifs = append(transformedNotifs, n)
		}

		returnedNotifs = transformedNotifs

		templates := &promptui.SelectTemplates{
			Label:    "{{ .Sender.Name }}",
			Active:   "-> [{{ .Sender.Name | cyan }}] ",
			Inactive: "   [{{ .Sender.Name | cyan }}] ",
			Details: `
--------- Details ----------
{{ "Sender:" | faint }}	{{ .Sender.Name }}: {{ .content }}`,
		}

		// searcher := func(input string, index int) bool {
		// 	notif := returnedNotifs[index]
		// 	name := strings.Replace(strings.ToLower(notif.Comment.Content), " ", "", -1)
		// 	input = strings.Replace(strings.ToLower(input), " ", "", -1)

		// 	return strings.Contains(name, input)
		// }

		searcher := func(input string, index int) bool {
			notif := returnedNotifs[index]
			name := strings.Replace(strings.ToLower(notif.content), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		}

		prompt := promptui.Select{
			Label:     "Select notification",
			Items:     returnedNotifs,
			Templates: templates,
			Size:      6,
			Searcher:  searcher,
		}

		i, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Printf("You choose number %d: %s\n", i+1, returnedNotifs[i].Sender.Name)

	},
}

func init() {
	// notifCmd.Flags().StringVarP(&BaseBranch, "branch", "b", "master", "Designate a branch (other than master) to merge to.")
	RootCmd.AddCommand(notifCmd)
}
