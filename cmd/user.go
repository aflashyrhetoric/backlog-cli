package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aflashyrhetoric/backlog-cli/utils"

	"github.com/spf13/cobra"
)

// User .. represents the current Backlog user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"mailAddress"`
	Language string `json:"lang"`
}

// CurrentUser ..
type CurrentUser User

var userCmd = &cobra.Command{
	Hidden: true,
	Use:    "user",
	Short:  "Returns data about myself",
	Run: func(cmd *cobra.Command, args []string) {

		endpoint := UserSelfEndpoint()

		// Unmarshal response data to variable
		var currentUser User
		responseData := utils.Get(endpoint)
		json.Unmarshal(responseData, &currentUser)

		fmt.Printf("User info for: %v", currentUser)
	},
}

// GetCurrentUser .. Returns the current user
func GetCurrentUser() User {

	endpoint := UserEndpoint()
	responseData := utils.Get(endpoint)

	var returnedUser User
	json.Unmarshal(responseData, &returnedUser)

	return returnedUser
}

// GetUserList ... returns a list of users in your space
func GetUserList() []User {
	endpoint := UserEndpoint()

	responseData := utils.Get(endpoint)

	var returnedUsers []User
	json.Unmarshal(responseData, &returnedUsers)

	return returnedUsers
}

func (u *User) printUserTable() {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Email: %s\n", u.Email)
	fmt.Printf("Link to Profile: %s/user/%s\n", GlobalConfig.BaseURL, u.Username)
	fmt.Printf("Link to Gantt Chart: %s/user/%s#usergantt\n", GlobalConfig.BaseURL, u.Username)
	fmt.Printf("Link to Settings Page: %s/EditProfile.action\n", GlobalConfig.BaseURL)
}

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get quick links for the current user",

	Run: func(cmd *cobra.Command, args []string) {
		u := GlobalConfig.User
		u.printUserTable()
	},
}

func init() {
	RootCmd.AddCommand(meCmd)
	RootCmd.AddCommand(userCmd)
}
