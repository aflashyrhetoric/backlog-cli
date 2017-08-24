package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var httpClient *http.Client

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "A brief description of your command",
	Long:  `to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		apiUrl := "/api/v2/users/myself"
		endpoint := Endpoint(apiUrl)

		// Fetch
		response, err := http.Get(endpoint)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		// Response
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))

		// A Response struct to map the Entire Response
		type User struct {
			Name     string `json:"name"`
			Email    string `json:"mailAddress"`
			Username string `json:"nulabAccount.uniqueId"`
		}

		var responseObject User

		json.Unmarshal(responseData, &responseObject)
		fmt.Println(responseObject.Name)
		fmt.Println(responseObject.Email)
	},
}

func init() {
	RootCmd.AddCommand(meCmd)
}
