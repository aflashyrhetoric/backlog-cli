package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

// GetCurrentRepo .. Returns current repository
func GetCurrentRepo() *git.Repository {
	var path string
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(path)
	ErrorCheck(err)

	return repo
}

// GetCurrentRepositoryName ... returns current repository name
func GetCurrentRepositoryName() string {

	//  FIXME: Get repository name through receiver function later
	repo := GlobalConfig.Repository

	// Open the 'origin' remote
	originRemote, err := repo.Remote("origin")
	ErrorCheck(err)

	// Fetch references from 'origin'
	repoReferences := originRemote.String()

	// Capture repository name from reference
	reg := regexp.MustCompile(`\/[A-Z]*\/(.*)\.git`)
	repositoryCapturedString := reg.FindSubmatch([]byte(repoReferences))
	repositoryReferenceName := string(repositoryCapturedString[1])

	return repositoryReferenceName
}

// GetCurrentBranch .. Gets current branch name.
func GetCurrentBranch() string {
	var path string
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(path)
	ErrorCheck(err)

	branchName, err := repo.Head()
	ErrorCheck(err)

	CurrentBranchName := branchName.Name()[11:]
	ErrorCheck(err)

	return string(CurrentBranchName)
}

// CheckIfBacklogRepo ... checks if the current repository is for Backlog (or other, e.g Github)
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
