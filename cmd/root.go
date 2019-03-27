package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// git "github.com/src-d/go-git"
	git "gopkg.in/src-d/go-git.v4"

	//"gopkg.in/src-d/go-git.v4/plumbing"
	"net/http"
	"os"
)

var configFile string
var hc = http.Client{}

// Hard-code path string
var path = "/Users/kevinoh/Nulab/cacoo-blog"
var formContentType = "Content-Type:application/x-www-form-urlencoded"

// RootCmd ... The primary main cobra command
var RootCmd = &cobra.Command{
	Use:   "backlog-cli",
	Short: "Use Backlog from the command line.",
	Long:  `Use Backlog from the command line to create pull requests, check issue status, access web pages, etc.`,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute ... runs the command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if configFile != "" {
		fmt.Println("Config found. Loading...")
		viper.SetConfigFile(configFile)
	} else {
		fmt.Println("Config not found. Setting defaults...")

		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")

		// read in environment variables that match
		viper.AutomaticEnv()
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

// FIXME: Temporary getter for Project Key

// Prints out a []byte response
func printResponse(responseData []byte) {
	fmt.Println(string(responseData[:]))
}

// Endpoint returns an endpoint
func Endpoint(apiURL string) string {
	// FIXME: We should just take SpaceID and build the "baseURL" from that
	baseURL := viper.GetString("BASEURL")
	key := "?apiKey=" + viper.GetString("API_KEY")
	endpoint := baseURL + apiURL + key
	return endpoint
}

// Checks for errors
func errorCheck(err error) {
	if err != nil {
		fmt.Printf("#%v", err)
	}
}

// Checks for errors, panics if found
func errorPanic(err error) {
	if err != nil {
		fmt.Printf("#%v", err)
		panic(err)
	}
}

// Gets current branch name.
func currentBranch(path string) string {
	repo, err := git.PlainOpen(path)
	errorCheck(err)

	branchName, err := repo.Head()

	/* Fetch the branch name by splicing the string
	 * FIXME: IS there a more reliable way?
	 */
	currentBranchName := branchName.Name()[11:]
	errorCheck(err)

	return string(currentBranchName)
}
