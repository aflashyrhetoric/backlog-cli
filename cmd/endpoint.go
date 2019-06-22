package cmd

import "fmt"

func endpoint(apiURL string) string {
	endpoint := fmt.Sprintf("%s/api/%s/%s?apiKey=%s", GlobalConfig.BaseURL, APIVersion(), apiURL, GlobalConfig.APIKey)
	return endpoint
}

func APIVersion() string {
	return fmt.Sprintf("v%d", GlobalConfig.BacklogAPIVersion)
}

// UserEndpoint ... Endpoint for querying users
func UserEndpoint() string {
	return endpoint("users")
}

// UserSelfEndpoint ... Endpoint for querying users
func UserSelfEndpoint() string {
	return endpoint("users/myself")
}

// PREndpoint ... Endpoint for querying pull requests
func PREndpoint() string {
	apiURL := fmt.Sprintf("projects/%s/git/repositories/%s/pullRequests", GlobalConfig.ProjectKey, GlobalConfig.RepositoryName)
	return endpoint(apiURL)
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
