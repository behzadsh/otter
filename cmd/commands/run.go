package commands

import "github.com/spf13/cobra"

var runCommand = &cobra.Command{
	Use:     "run",
	Short:   "Run all tests scenario",
	Example: "otter run",
	Args:    cobra.NoArgs,
	RunE:    runRunCommand,
}

var ()

func runRunCommand(cmd *cobra.Command, args []string) error {
	return nil
}
