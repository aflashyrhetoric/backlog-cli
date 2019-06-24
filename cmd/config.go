package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/yaml.v2"
)

// GlobalConfig .. the global configuration for backlog-cli
var GlobalConfig Config

// Config .. the configuration struct for backlog-cli
type Config struct {
	User              User            `yaml:",omitempty"`
	BaseURL           string          `yaml:"BASEURL,omitempty"`
	APIKey            string          `yaml:"API_KEY"`
	DebugMode         bool            `yaml:"DEBUG_MODE"`
	ProjectKey        string          `yaml:",omitempty"`
	Repository        *git.Repository `yaml:",omitempty"`
	RepositoryName    string          `yaml:",omitempty"`
	CurrentBranch     string          `yaml:",omitempty"`
	CurrentIssue      Issue           `yaml:",omitempty"`
	BacklogAPIVersion int             `yaml:",omitempty"`
}

// var initCmd = &cobra.Command{
// 	Use:   "init",
// 	Short: "Set Backlog config through an interactive prompt.",
// 	Run: func(cmd *cobra.Command, args []string) {

// 		InitialSetup()

// 	},
// }

// InitialSetup ... receives prompts from the user to set up backlog-config.yaml
func InitialSetup() {

	// First, need BASEURL
	baseURL, err := PromptInput("Enter BaseURL", "https://yourspace.backlog.com")

	// Remove whitespace and trailing slash
	baseURL = strings.TrimSpace(baseURL)
	baseURL = strings.TrimRight(baseURL, "/")

	// Get Link to Page
	apiPageLink := LinkToAPIPage(baseURL)
	apiKey, err := PromptInput("Enter your API Key", fmt.Sprintf("Visit %s and create an API Key for Backlog CLI, and enter it below", apiPageLink))
	ErrorCheck(err)
	apiKey = strings.TrimSpace(apiKey)

	debugMode, err := PromptConfirm("Enable debug/developer mode for more developer information?")
	ErrorCheck(err)

	// Get home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configData := Config{
		BaseURL:   baseURL,
		APIKey:    apiKey,
		DebugMode: debugMode,
	}
	configByteData, err := yaml.Marshal(&configData)
	ErrorCheck(err)

	homeDir := usr.HomeDir
	configPath := fmt.Sprintf("%s/backlog-config.yaml", homeDir)

	err = ioutil.WriteFile(configPath, configByteData, 0644)
	ErrorPanic(err)

	fmt.Println("Backlog CLI configuration complete.")

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

// func init() {
// 	RootCmd.AddCommand(initCmd)
// }
