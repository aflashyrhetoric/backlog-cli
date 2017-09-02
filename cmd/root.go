package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	git "gopkg.in/src-d/go-git.v4"
	//"gopkg.in/src-d/go-git.v4/plumbing"
	"net/http"
	"os"
)

var cfgFile string
var hc = http.Client{}
var path string = "/home/wdkevo/go/src/backlogtool.com/BLGTEST/testrepo"
var formContentType string = "Content-Type:application/x-www-form-urlencoded"

var RootCmd = &cobra.Command{
	Use:   "backlog-cli",
	Short: "Use Backlog from the command line.",
	Long:  `to quickly create a Cobra application.`,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// FIXME: Temporary getter for Project Key
func ProjectKey() string {
	return viper.GetString("PROJECT_KEY")
}

// FIXME: Temporary getter for repository name
func Repo() string {
	return viper.GetString("REPOSITORY_NAME")
}

func initConfig() {
	if cfgFile != "" {
		fmt.Println("Config found. Loading...")
		viper.SetConfigFile(cfgFile)
	} else {
		fmt.Println("Config not found. Setting defaults...")

		viper.SetConfigName(".backlog-cli")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		// viper.SetEnvPrefix("backlog")

		// read in environment variables that match
		viper.AutomaticEnv()
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

// Prints out a []byte response
func printResponse(responseData []byte) {
	fmt.Println(string(responseData[:]))
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
	currentBranchName := branchName.Name()[11:]
	errorCheck(err)

	return string(currentBranchName)
}
