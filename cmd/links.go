package cmd

import (
	"fmt"
)

// LinkToPRPage ... returns a hyperlink to the PR based on its ID
func LinkToPRPage(n int) string {
	return fmt.Sprintf("%s/git/%s/%s/pullRequests/%d", GlobalConfig.BaseURL, GlobalConfig.ProjectKey, GlobalConfig.RepositoryName, n)
}

func LinkToIssuePage(issueKey string) string {
	return fmt.Sprintf("%s/view/%s", issueKey)
}
