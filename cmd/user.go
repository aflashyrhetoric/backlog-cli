package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aflashyrhetoric/backlog-cli/utils"

	"github.com/spf13/cobra"
)

// User .. represents a Backlog user
type User struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	Language    string `json:"lang"`
	MailAddress string `json:"mailAddress"`
}

// Gets
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Returns data about myself",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		// Set up Endpoint
		apiURL := "/api/v2/users/myself"
		endpoint := Endpoint(apiURL)

		// Unmarshal response data to variable
		var currentUser User
		responseData := utils.Get(endpoint)
		json.Unmarshal(responseData, &currentUser)
		fmt.Printf("User info for: %v", currentUser)
	},
}

func init() {
	RootCmd.AddCommand(userCmd)
}
