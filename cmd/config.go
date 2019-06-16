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
}

// Helper to set the User to the GlobalConfig
func (c *Config) setUser(user User) {
	c.User = user
}

func (c *Config) setBranch(branch string) {
	c.CurrentBranch = branch
}

func (c *Config) setIssue(issue Issue) {
	c.CurrentIssue = issue
}
