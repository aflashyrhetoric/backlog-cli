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

		// By default, get Issue ID from current branch name if possible
		apiURL := "/api/v2/users/myself"
		endpoint := Endpoint(apiURL)
		responseData := utils.Get(endpoint)
		var currentUser User
		json.Unmarshal(responseData, &currentUser)
		fmt.Printf("User info for: %v", currentUser)
	},
}

func init() {
	RootCmd.AddCommand(userCmd)
}
