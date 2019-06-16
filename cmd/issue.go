package cmd

import (
	"encoding/json"
	"strconv"

	"github.com/aflashyrhetoric/backlog-cli/utils"
)

// Issue .. the configuration struct for backlog-cli
type Issue struct {
	ID          int    `json:"id"`
	IssueKey    string `json:"issueKey"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

// Key .. Returns the issue's key (e.g. "BLG-123") as a string
func (i *Issue) Key() string {
	apiURL := "/api/v2/issues/" + strconv.Itoa(i.ID)
	endpoint := Endpoint(apiURL)
	responseData := utils.Get(endpoint)

	// A Response struct to map the Entire Response
	var returnedIssue Issue
	json.Unmarshal(responseData, &returnedIssue)

	return returnedIssue.IssueKey
}

// Desc .. Returns the description as a string
// func (i *Issue) Desc() string {
// 	apiURL := "/api/v2/issues/" + strconv.Itoa(i.ID)
// 	endpoint := Endpoint(apiURL)
// 	responseData := utils.Get(endpoint)

// 	// A Response struct to map the Entire Response
// 	var returnedIssue Issue
// 	json.Unmarshal(responseData, &returnedIssue)

// 	return returnedIssue.IssueKey
// }

// // DescBranchName .. Returns a truncated version of the description as a branch name
// func (i *Issue) DescBranchName() string {
// 	issueDescription := i.Desc()

// 	return returnedIssue.IssueKey
// }
