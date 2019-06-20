package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aflashyrhetoric/backlog-cli/utils"
	e "github.com/kyokomi/emoji"
	a "github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"

	"github.com/spf13/cobra"
)

// Whether or not the user wants to open the
var interactiveFlag bool
var tableViewFlag bool

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

func (n *notificationSender) string() string {
	return n.Name
}

type notificationComment struct {
	Content string `json:"content"`
}
type notificationPullRequestComment struct {
	Content string `json:"content"`
}

var notifCmd = &cobra.Command{
	Use:     "notification",
	Aliases: []string{"notifications", "n"},
	Short:   "Read and open your notifications",
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

		var truncatedNotifications []Notification

		for _, n := range returnedNotifs {
			n.truncateName()
			truncatedNotifications = append(truncatedNotifications, n)
		}

		if !interactiveFlag {
			if tableViewFlag {
				printNotificationTable(truncatedNotifications)
			} else {
				listNotifications("bell", "Notifications", truncatedNotifications)
			}
		} else {

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
				Items:     truncatedNotifications,
				Templates: templates,
				Size:      6,
			}

			i, _, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			notificationURL := fmt.Sprintf("%s/globalbar/notifications/redirect/%d\n", GlobalConfig.BaseURL, returnedNotifs[i].ID)
			openBrowser(notificationURL)
		}
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
	case 4:
		text = n.Sender.Name + " updated an issue."
		break
	case 5:
		text = n.Sender.Name + " attached a file."
		break
	case 9:
		text = "Issue type unknown"
		break
	case 10:
		text = n.PullRequestComment.Content
		break
	case 11:
		text = n.PullRequestComment.Content
		break
	case 12:
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

func listNotifications(emoji string, message string, notifList []Notification) {
	e.Print("\n\n:" + emoji + ":")
	fmt.Println(a.White("[" + message + "]").BgBrightBlack())
	for i, notif := range notifList {
		fmt.Printf("[%d] %s: %s\n", i, notif.Sender, notif.Content)
	}
}

func printNotificationTable(notifList []Notification) {
	var data [][]string

	for _, notif := range notifList {
		data = append(data, []string{
			notif.Sender.string(),
			notif.Content,
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Sender", "Content"})

	for _, v := range data {
		table.Append(v)
	}
	// Send output
	table.Render()
}

func init() {
	notifCmd.Flags().BoolVarP(&interactiveFlag, "interactive", "i", false, "Include to open interactive notification viewer")
	notifCmd.Flags().BoolVarP(&tableViewFlag, "table", "t", false, "Include to view notifications as a table")
	RootCmd.AddCommand(notifCmd)
}
