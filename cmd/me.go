package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/dghubble/sling"
)

// meCmd represents the me command
var meCmd = &cobra.Command{
	Use:   "me",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		baseurl:= viper.GetString("API_KEY")
		url:= "/api/v2/users/myself"
		key:= '?apiKey=' + viper.GetString("API_KEY") 
		endpoint := baseurl + url + key +

		params := &Params{Count: 5}

		req, err := sling.New().Get(endpoint).QueryStruct().Request()
		client.Do(req)
	},
}

func init() {
	RootCmd.AddCommand(meCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// meCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// meCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
