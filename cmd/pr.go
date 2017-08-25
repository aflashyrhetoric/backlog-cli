package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates a Backlog Pull Request for the current branch",
	Long:  `abc`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := git.PlainOpen(path)
		if err != nil {
			fmt.Printf("#%v", err)
			panic(err)
		}
		branches, err := repo.Branches()
		if err != nil {
			panic(err)
		}

		err = branches.ForEach(reference)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(prCmd)
}

func reference(refer *plumbing.Reference) error {
	//fmt.Printf("%#v\n", refer)
	fmt.Printf("%s\n", refer.Name())
	return nil
}
