package cmd

import (
"fmt"
"net/http"

"github.com/spf13/cobra"
"github.com/spf13/viper"
)

var httpClient *http.Client

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "A brief description of your command",
	Long: `to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		apiUrl:= "/api/v2/users/myself"
		endpoint:= Endpoint(apiUrl)

		fmt.Println(endpoint)
	},
}

func init() {
	RootCmd.AddCommand(meCmd)
	RootCmd.Prep()
}