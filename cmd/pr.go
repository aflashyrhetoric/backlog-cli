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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := git.PlainOpen("/Users/yen-chiehchen/workspace/src/bowery-golang_demo")
		if err != nil {
			fmt.Printf("#%v", err)
			panic(err)
		}

		branches, err := repo.Branches()
		if err != nil {
			panic(err)
		}

		err = branches.ForEach(referrals)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(prCmd)
}

func referrals(refer *plumbing.Reference) error {
	fmt.Printf("%#v\n", refer)
	fmt.Printf("%s\n", refer.Name())

	return nil
}
