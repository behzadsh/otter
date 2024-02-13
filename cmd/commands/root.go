package commands

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/behzadsh/otter/internal/configs"
)

var rootCmd = &cobra.Command{
	Use:   "otter",
	Short: "Otter is a lightweight and user-friendly API testing tool",
	// Version: "",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	SilenceErrors: silenceErrors(),
}

func init() {
	rootCmd.SetErrPrefix("otter:")

	rootCmd.AddCommand(initCommand)
	rootCmd.AddCommand(newCommand)
}

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}

func silenceErrors() bool {
	return configs.DevMode()
}
