package cmd

import "fmt"

// StringBuilder ... builds endpoints for querying
type StringBuilder struct {
	BaseURL    string
	ProjectKey string
	IssueKey   string
	APIKey     string
	APIVersion string
}

// NewStringBuilder ... Returns new StringBuilder for easy querying
func NewStringBuilder() *StringBuilder {
	return &StringBuilder{
		BaseURL:    GlobalConfig.BaseURL,
		ProjectKey: GlobalConfig.ProjectKey,
		APIKey:     GlobalConfig.APIKey,
		APIVersion: "v2",
	}
}

func (s *StringBuilder) endpoint(apiURL string) string {
	endpoint := fmt.Sprintf("%s/api/%s/%s?apiKey=%s", s.BaseURL, s.APIVersion, apiURL, s.APIKey)
	return endpoint
}

// UserEndpoint ... Endpoint for querying users
func (s *StringBuilder) UserEndpoint() string {
	return s.endpoint("users/myself")
}

// PREndpoint ... Endpoint for querying pull requests
func (s *StringBuilder) PREndpoint() string {
	apiURL := fmt.Sprintf("projects/%s/git/repositories/%s/pullRequests", GlobalConfig.ProjectKey, GlobalConfig.RepositoryName)
	return s.endpoint(apiURL)
}
