package cmd

import (
	"encoding/json"
	"strconv"

	"github.com/aflashyrhetoric/backlog-cli/utils"
)

// Issue .. the configuration struct for backlog-cli
type Issue struct {
	ID       int    `json:"id"`
	IssueKey string `json:"issueKey"`
}

// Key .. Givens an Issue's PrimaryKey ID, returns the issue as a string
func (i *Issue) Key() string {
	apiURL := "/api/v2/issues/" + strconv.Itoa(i.ID)
	endpoint := Endpoint(apiURL)
	responseData := utils.Get(endpoint)

	// A Response struct to map the Entire Response
	var returnedIssue Issue
	json.Unmarshal(responseData, &returnedIssue)

	return returnedIssue.IssueKey
}
