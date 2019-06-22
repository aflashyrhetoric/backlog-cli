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

		apiURL := "/api/v2/users/myself"
		endpoint := Endpoint(apiURL)

		// Unmarshal response data to variable
		var currentUser User
		responseData := utils.Get(endpoint)
		json.Unmarshal(responseData, &currentUser)
		fmt.Printf("User info for: %v", currentUser)
	},
}

// GetCurrentUser .. Returns the current user
func GetCurrentUser() User {
	apiURL := "/api/v2/users/myself"
	endpoint := Endpoint(apiURL)

	responseData := utils.Get(endpoint)

	var returnedUser User
	json.Unmarshal(responseData, &returnedUser)

	return returnedUser
}

func init() {
	RootCmd.AddCommand(userCmd)
}
