package cmd

import (
	"backlog-cli/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	//"gopkg.in/src-d/go-git.v4/plumbing"
	"net/url"
	"regexp"
	"strconv"
)

// Gets
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates a Backlog Pull Request for the current branch",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// ---------------------------------------------------------

		// By default, get Issue ID from current branch name if possible
		currentBranch := currentBranch(path)
		reg := regexp.MustCompile("([a-zA-Z]+-[0-9]*)")
		issueId := string(reg.Find([]byte(currentBranch)))
		//fmt.Println(issueId)

		apiUrl := "issues/" + string(issueId)

		if currentBranch == "staging" || currentBranch == "dev" || currentBranch == "beta" {
			fmt.Printf("You're currently on the %v branch. Please switch to an issue branch and try again.", issueId)
		} else if currentBranch == "0" {
			fmt.Println("Invalid branch. Try again.")
		} else {
			fmt.Printf("Creating PR for %v branch.", currentBranch)
		}

		endpoint := utils.Endpoint(apiUrl)
		type Issue struct {
			Id int `json:"id"`
		}
		responseData := utils.Get(endpoint)
		var currentIssue Issue
		json.Unmarshal(responseData, &currentIssue)
		// Convert integer -> string for use in later functions
		issueId = strconv.Itoa(currentIssue.Id)

		//fmt.Println(issueId)

		// Create the form, request, and send the POST request
		// ---------------------------------------------------------
		p := ProjectKey()
		r := Repo()
		apiUrl = "projects/" + p + "/git/repositories/" + r + "/pullRequests"

		//apiUrl = "test"
		endpoint = utils.Endpoint(apiUrl)

		// Build out Form
		form := url.Values{}
		form.Add("summary", "Test summary")
		form.Add("description", "Test description")
		form.Add("base", "master")
		form.Add("branch", currentBranch)
		form.Add("issueId", issueId)

		responseData = utils.Post(endpoint, form)

		// A Response struct to map the Entire Response
		type PullRequest struct {
			Summary     string `json:"summary"`
			Description string `json:"description"`
			Base        string `json:"base"`
			Branch      string `json:"branch"`
			Id          int    `json:"id"`
		}
		var returnedPullRequestCount PullRequest
		json.Unmarshal(responseData, &returnedPullRequestCount)
		printResponse(responseData)

		fmt.Println(returnedPullRequestCount.Summary)

		//err = branches.ForEach(reference)
		//errorCheck(err)
	},
}

func init() {
	RootCmd.AddCommand(prCmd)
}

//func reference(refer *plumbing.Reference) error {
//	//fmt.Printf("%#v\n", refer)
//	fmt.Printf("%s\n", refer.Name()[11:])
//	return nil
//}
