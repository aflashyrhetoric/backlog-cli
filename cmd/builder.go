package cmd

import (
	"fmt"

	e "github.com/kyokomi/emoji"
)

var SB StringBuilder

// StringBuilder ... builds endpoints for querying
type StringBuilder struct {
	BaseURL        string
	ProjectKey     string
	RepositoryName string
	IssueKey       string
	APIKey         string
	APIVersion     string
}

// NewStringBuilder ... Returns new StringBuilder for easy querying
// func NewStringBuilder() *StringBuilder {
// 	return &StringBuilder{
// 		BaseURL:    GlobalConfig.BaseURL,
// 		ProjectKey: GlobalConfig.ProjectKey,
// 		APIKey:     GlobalConfig.APIKey,
// 		APIVersion: "v2",
// 	}
// }

func endpoint(apiURL string) string {
	endpoint := fmt.Sprintf("%s/api/%s/%s?apiKey=%s", GlobalConfig.BaseURL, APIVersion(), apiURL, GlobalConfig.APIKey)
	return endpoint
}

func APIVersion() string {
	return fmt.Sprintf("v%d", GlobalConfig.BacklogAPIVersion)
}

// UserEndpoint ... Endpoint for querying users
func UserEndpoint() string {
	return endpoint("users/myself")
}

// PREndpoint ... Endpoint for querying pull requests
func PREndpoint() string {
	apiURL := fmt.Sprintf("projects/%s/git/repositories/%s/pullRequests", GlobalConfig.ProjectKey, GlobalConfig.RepositoryName)
	return endpoint(apiURL)
}

// LinkToPRPage ... returns a hyperlink to the PR based on its ID
func LinkToPRPage(n int) string {
	return fmt.Sprintf("%s/git/%s/%s/pullRequests/%d", GlobalConfig.BaseURL, GlobalConfig.ProjectKey, GlobalConfig.RepositoryName, n)
}

// IssueListEndpoint ... Endpoint for querying issues list
func IssueListEndpoint() string {
	return endpoint("issues")
}

// IssueEndpoint ... Endpoint for querying specific issue ID
func IssueEndpoint(issueID string) string {
	apiURL := fmt.Sprintf("issues/%s", issueID)
	return endpoint(apiURL)
}

// NotificationEndpoint ... Endpoint for querying specific notifications
func NotificationEndpoint() string {
	return endpoint("notifications")
}

// EmojiPrefixMessage ... LogGlobalConfig...an emoji....
func EmojiPrefixMessage(emoji string) {
	e.Print("\n\n:" + emoji + ":")
}
