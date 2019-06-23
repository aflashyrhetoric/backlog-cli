package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"net/http"
	"os"
)

var configFile string
var hc = http.Client{}

// RootCmd ... The primary main cobra command
var RootCmd = &cobra.Command{
	Use:   "backlog-cli",
	Short: "Use Backlog from the command line.",
	Long:  `Use Backlog from the command line to create pull requests, check issue status, access web pages, etc.`,
}

// Execute ... runs the command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	CheckIfBacklogRepo()
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

	if configFile != "" {
		debugPrint("Using config file: %s\n", configFile)

		GlobalConfig = Config{
			BaseURL:           viper.GetString("BASEURL"),
			APIKey:            viper.GetString("API_KEY"),
			Repository:        GetCurrentRepo(),
			CurrentBranch:     GetCurrentBranch(),
			BacklogAPIVersion: 2,
		}

		// SB := NewStringBuilder()

		// Configuration that requires HTTP, call them after GlobalConfig is initialized
		GlobalConfig.RepositoryName = GetCurrentRepositoryName()
		GlobalConfig.ProjectKey = GetProjectKey()
		GlobalConfig.CurrentIssue = GetCurrentIssue()
		GlobalConfig.User = GetCurrentUser()

	} else {
		fmt.Println("Config file not found. Initializing setup...")
		InitialSetup()
	}

	viper.AutomaticEnv()
}
