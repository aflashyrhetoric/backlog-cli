package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get quick links for the current user",

	Run: func(cmd *cobra.Command, args []string) {
		currentUser := GetCurrentUser()
		fmt.Printf("Name: %s\n", currentUser.Name)
		fmt.Printf("Email: %s\n", currentUser.Email)
		fmt.Printf("Link to Profile: %s/user/%s\n", GlobalConfig.BaseURL, currentUser.Username)
		fmt.Printf("Link to Gantt Chart: %s/user/%s#usergantt\n", GlobalConfig.BaseURL, currentUser.Username)
		fmt.Printf("Link to Settings Page: %s/EditProfile.action\n", GlobalConfig.BaseURL)
	},
}

func init() {
	RootCmd.AddCommand(meCmd)
}
