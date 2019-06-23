package cmd

import (
	"fmt"
	"log"
	"os/user"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

// GlobalConfig .. the global configuration for backlog-cli
var GlobalConfig Config

// Config .. the configuration struct for backlog-cli
type Config struct {
	User              User
	BaseURL           string
	APIKey            string
	ProjectKey        string
	Repository        *git.Repository
	RepositoryName    string
	CurrentBranch     string
	CurrentIssue      Issue
	DebugMode         bool
	BacklogAPIVersion int
}

// InitialSetup ... receives prompts from the user to set up backlog-config.yaml
func InitialSetup() {

	// First, need BASEURL
	baseURL, err := PromptInput("What is your BASEURL?")

	// Remove whitespace and trailing slash
	baseURL = strings.TrimSpace(baseURL)
	baseURL = strings.TrimRight(baseURL, "/")

	// Get Link to Page
	apiPageLink := LinkToAPIPage(baseURL)
	apiKey, err := PromptInput(fmt.Sprintf("Visit %s and create an API Key for Backlog CLI, and enter it below", apiPageLink))
	ErrorCheck(err)
	apiKey = strings.TrimSpace(apiKey)

	debugMode, err := PromptConfirm("Enable debug/developer mode for more developer information?")
	ErrorCheck(err)

	// Get home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDir := usr.HomeDir
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
