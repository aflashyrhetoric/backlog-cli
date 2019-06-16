package cmd

import git "gopkg.in/src-d/go-git.v4"

// GlobalConfig .. the global configuration for backlog-cli
var GlobalConfig Config

// Config .. the configuration struct for backlog-cli
type Config struct {
	User           User
	BaseURL        string
	APIKey         string
	ProjectKey     string
	Repository     *git.Repository
	RepositoryName string
	CurrentBranch  string
	CurrentIssue   Issue
	DebugMode      bool
}
