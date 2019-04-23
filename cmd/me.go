package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aflashyrhetoric/backlog-cli/utils"

	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "A brief description of your command",
	Long:  `to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		apiURL := "/api/v2/users/myself"
		endpoint := Endpoint(apiURL)

		// Fetch
		responseData := utils.Get(endpoint)

		// A Response struct to map the Entire Response
		type User struct {
			Name  string `json:"name"`
			Email string `json:"mailAddress"`
			ID    string `json:"userId"`
		}

		var returnedUser User

		json.Unmarshal(responseData, &returnedUser)
		fmt.Println(returnedUser.Name)
		fmt.Println(returnedUser.Email)
	},
}

func init() {
	RootCmd.AddCommand(meCmd)
}
