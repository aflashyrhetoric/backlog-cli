package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aflashyrhetoric/backlog-cli/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get quick links for the current user",
	Long:  `to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		apiURL := "/api/v2/users/myself"
		endpoint := Endpoint(apiURL)

		// Fetch
		responseData := utils.Get(endpoint)

		// A Response struct to map the Entire Response
		type User struct {
			Name     string `json:"name"`
			Email    string `json:"mailAddress"`
			Language string `json:"lang"`
			ID       string `json:"userId"`
		}

		var returnedUser User

		json.Unmarshal(responseData, &returnedUser)
		fmt.Printf("Name: %s\n", returnedUser.Name)
		fmt.Printf("Email: %s\n", returnedUser.Email)
		fmt.Printf("Link to Profile: %s/user/%s\n", viper.GetString("BASEURL"), returnedUser.ID)
		fmt.Printf("Link to Gantt Chart: %s/user/%s#usergantt\n", viper.GetString("BASEURL"), returnedUser.ID)
	},
}

func init() {
	RootCmd.AddCommand(meCmd)
}
