package cmd

import (
	"fmt"
	"net/http"
	// "net/url"
	// "encoding/json"
	"io/ioutil"
	// "strings"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var httpClient *http.Client

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		apiUrl:= "/api/v2/users/myself"
		endpoint:= Endpoint(apiUrl)

		resp, err := http.Get(endpoint)
		if err != nil {
			fmt.Println("There's been a fatal error.")
		}
		
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		fmt.Println(body)
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
