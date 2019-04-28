package cmd

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"net/http"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

var configFile string
var hc = http.Client{}

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
	viper.SetConfigName("backlog-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")

	// Load config
	err := viper.ReadInConfig()
	ErrorCheck(err)

	// Set name
	configFile = viper.ConfigFileUsed()
	viper.SetConfigName(configFile)

	// FIXME: Remove debug info for production build
	if configFile != "" {
		fmt.Printf("Config found. Loaded %s\n", configFile)

		GlobalConfig = Config{
			BaseURL:        viper.GetString("BASEURL"),
			APIKey:         viper.GetString("API_KEY"),
			ProjectKey:     ProjectKey(),
			Repository:     Repository(),
			RepositoryName: RepositoryName(),
			CurrentBranch:  CurrentBranch(),
		}

		GetCurrentUser()

	} else {
		fmt.Println("Config not found. Please create a config at $HOME/backlog-config.yaml")
	}

	viper.AutomaticEnv()
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

// Repository .. Returns current repository
func Repository() *git.Repository {
	var path string
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(path)
	ErrorCheck(err)

	return repo
}

// ProjectKey ... Returns the project key for the configuration
func ProjectKey() string {
	var cb string
	cb = CurrentBranch()
	cb = strings.ToLower(cb)
	if cb == "staging" || cb == "dev" || cb == "develop" || cb == "beta" {
		path, err := os.Getwd()
		ErrorCheck(err)
		path = strings.ToLower(path)
		// TODO: FETCH LIST OF PROJECTS FROM BACKLOG AND USE REGULAR EXPRESSIONS AND STRING SIMILARITY ALGORITHM TO MATCH CURRENT WORKING DIRECTORY TO PROJECTKEY
		// RESTRUCTURE THE PROJECTKEY() FUNCTION TO MAKE IT CLEAR THAT IT USES A CASCADE TO TRY AND DETECT AN ACCURATE KEY
		// IF ALL ELSE FAILS OFFER A CONFIG-BASED WAY (IN THE YAML FILE) TO SET UP DIRECTORY-BASED PROJECT KEY MATCHING (SEEMS BEST WAY LONG-TERM)
		reg := regexp.MustCompile(`(.*)-[0-9]+`)
		projectKeyCapturedString := reg.FindSubmatch([]byte(cb))
		projectKeyReferenceName := string(projectKeyCapturedString[1])
		return projectKeyReferenceName
	}
	reg := regexp.MustCompile(`(.*)-[0-9]+`)
	projectKeyCapturedString := reg.FindSubmatch([]byte(cb))
	projectKeyReferenceName := string(projectKeyCapturedString[1])
	return projectKeyReferenceName
}

// RepositoryName ... returns current repository name
func RepositoryName() string {

	repo := Repository()

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
