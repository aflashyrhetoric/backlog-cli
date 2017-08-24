package cmd

import (
"fmt"
"net/http"
"github.com/spf13/cobra"
// "github.com/libgit2/git2go"

"io/ioutil"
"log"
"os"
)

var httpClient *http.Client

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "A brief description of your command",
	Long: `to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		apiUrl:= "/api/v2/users/myself"
		endpoint:= Endpoint(apiUrl)

		response, err := http.Get(endpoint);
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}	

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))



	},
}

func init() {
	RootCmd.AddCommand(meCmd)
}
