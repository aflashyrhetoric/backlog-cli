package cmd

import (
	"fmt"
)

// LinkToPRPage ... returns a hyperlink to the PR based on its ID
func LinkToPRPage(n int) string {
	return fmt.Sprintf("%s/git/%s/%s/pullRequests/%d", GlobalConfig.BaseURL, GlobalConfig.ProjectKey, GlobalConfig.RepositoryName, n)
}

// LinkToIssuePage ... returns a hyperlink to the Current Issue Page
func LinkToIssuePage(issueKey string) string {
	return fmt.Sprintf("%s/view/%s", GlobalConfig.BaseURL, issueKey)
}

// LinkToAPIPage ... returns a hyperlink to the Current Issue Page
func LinkToAPIPage(baseURL string) string {
	return fmt.Sprintf("%s/EditApiSettings.action", baseURL)
}
