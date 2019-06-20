package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

// PrintResponse .. Prints out a []byte
func PrintResponse(responseData []byte) {
	fmt.Println(string(responseData[:]))
}

// Endpoint .. returns an endpoint
func Endpoint(apiURL string) string {
	// FIXME: We should just take SpaceID and build the "baseURL" from that
	endpoint := fmt.Sprintf("%s%s?apiKey=%s", GlobalConfig.BaseURL, apiURL, GlobalConfig.APIKey)
	return endpoint
}

// ErrorCheck .. Checks for error != nil
func ErrorCheck(err error) {
	if err != nil {
		fmt.Printf("#%v", err)
	}
}

// ErrorPanic .. Checks for errors, panics if found
func ErrorPanic(err error) {
	if err != nil {
		fmt.Printf("#%v", err)
		panic(err)
	}
}

// Truncate .. truncates a string to its max
func Truncate(s string) string {
	max := 45
	if len(s) >= max {
		return s[0:max]
	}
	return s
}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func CheckIfBacklogRepo() {
	var path string
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(path)
	ErrorCheck(err)

	// branchName, err := repo.Head()
	branchName, err := repo.Remote("origin")
	ErrorCheck(err)

	if !strings.Contains(branchName.String(), "git.backlog") {
		err = errors.New("this doesn't seem to be a Backlog repository - exiting")
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
