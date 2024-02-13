package commands

import "github.com/spf13/cobra"

var newCommand = &cobra.Command{
	Use:     "new [name]",
	Short:   "Create new test scenario file",
	Example: `otter new "user authentication"`,
	Args:    cobra.ExactArgs(1),
	RunE:    runNewCommand,
}

func runNewCommand(_ *cobra.Command, _ []string) error {
	return nil
}
