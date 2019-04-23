package cmd

import (
	"fmt"
	"log"
	"regexp"

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
	// fmt.Printf("Project Key '%v'\n", ProjectKey())
	// fmt.Printf("Repo name '%v'\n", RepoName())
	// fmt.Printf("branchname '%v'\n", CurrentBranch())
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")

	if configFile != "" {
		fmt.Println("Config found. Loading...")
		// FIXME Allow for dynamic configuration file loading
		viper.SetConfigName(configFile)
	} else {
		fmt.Println("Config not found. Setting defaults...")
		// read in environment variables that match
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			fmt.Println(err)
		}
		viper.AutomaticEnv()
	}
}

// CurrentBranch .. Gets current branch name.
func CurrentBranch() string {
	var path string
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(path)
	ErrorCheck(err)

	branchName, err := repo.Head()

	CurrentBranchName := branchName.Name()[11:]
	ErrorCheck(err)

	return string(CurrentBranchName)
}

// ProjectKey ... Returns the project key for the configuration
func ProjectKey() string {
	var cb string
	cb = CurrentBranch()
	reg := regexp.MustCompile(`(.*)-[0-9]+`)
	projectKeyCapturedString := reg.FindSubmatch([]byte(cb))
	projectKeyReferenceName := string(projectKeyCapturedString[1])
	return projectKeyReferenceName
}

// RepoName ... returns repository name
func RepoName() string {
	var path string
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Open Git repo
	repo, err := git.PlainOpen(path)
	ErrorCheck(err)

	// Open the 'origin' remote
	originRemote, err := repo.Remote("origin")
	ErrorCheck(err)

	// Fetch references from 'origin'
	repoReferences := originRemote.String()

	// Capture repository name from reference
	reg := regexp.MustCompile(`\/[A-Z]*\/(.*)\.git`)
	repositoryCapturedString := reg.FindSubmatch([]byte(repoReferences))
	repositoryReferenceName := string(repositoryCapturedString[1])

	return repositoryReferenceName
}
