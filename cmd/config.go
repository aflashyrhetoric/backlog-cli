package cmd

import git "gopkg.in/src-d/go-git.v4"

// Config .. represents the configuration for backlog-cli
var GlobalConfig Config

type Config struct {
	BaseURL        string
	ProjectKey     string
	Repository     *git.Repository
	RepositoryName string
	CurrentBranch  string
}
