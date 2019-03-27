package cmd

import "github.com/spf13/viper"

// Config .. represents the configuration for backlog-cli
type Config struct {
	projectKey     string
	repositoryName string
}

// ProjectKey ... Returns the project key for the configuration
func ProjectKey() string {
	return viper.GetString("PROJECT_KEY")
}

// FIXME: Temporary getter for repository name

// Repo ... returns repository name in viper
func Repo() string {
	return viper.GetString("REPOSITORY_NAME")
}
