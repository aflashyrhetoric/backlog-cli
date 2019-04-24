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
}

// Helper to set the User to the GlobalConfig
func (u *Config) setUser(user User) {
	u.User = user
}
