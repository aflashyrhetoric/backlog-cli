package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aflashyrhetoric/backlog-cli/utils"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

var myCmd = &cobra.Command{
	Hidden: true,
	Use:    "my",
	Short:  "Browse and open PULL REQUESTS and ISSUES that are assigned to you",
}

var myPullRequestCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List links to all associated pull requests",
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		p := GlobalConfig.ProjectKey
		r := GlobalConfig.RepositoryName
		// ! FIXME THIS ISNT THE RIGHT apiURL
		apiURL := "/api/v2/projects/" + p + "/git/repositories/" + r + "/pullRequests"

		//apiURL = "test"
		endpoint := Endpoint(apiURL)

		myPullRequests, err := getMyPullRequests(endpoint)
		ErrorCheck(err)

		if len(myPullRequests) == 0 {
			fmt.Println("There are no pull requests associated to this issue.")
		}
		if len(myPullRequests) > 0 {

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
				Items:     myPullRequests,
				Templates: templates,
				Size:      6,
			}

			i, _, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			// ! FIXME THIS ISNT THE RIGHT URL
			notificationURL := fmt.Sprintf("%s/globalbar/notifications/redirect/%d", GlobalConfig.BaseURL, myPullRequests[i])
			openBrowser(notificationURL)
		}

	},
}

func getMyPullRequests(endpoint string) ([]PullRequest, error) {

	// params for pull requests
	params := map[string]int{
		"assigneeId[]": GlobalConfig.User.ID,
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
	// prCreateCmd.Flags().StringVarP(&BaseBranch, "branch", "b", "master", "Designate a branch (other than master) to merge to.")
	// prCmd.AddCommand(prCreateCmd)
	// prCmd.AddCommand(prOpenCmd)
	// prCmd.AddCommand(prListCmd)
	RootCmd.AddCommand(myCmd)
}
