package commands

import (
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/behzadsh/otter/internal/configs"
)

var initCommand = &cobra.Command{
	Use:     "init",
	Short:   "Initiate the otter config file",
	Example: "otter init -u http://localhost:8000",
	Args:    cobra.NoArgs,
	RunE:    runInitCommand,
}

var (
	configFile    = ".otter.yml"
	baseURL       = "http://localhost"
	testCasesPath = "tests"
)

func init() {
	initCommand.Flags().StringVarP(&configFile, "config-path", "c", configFile, "define the path to the config file")
	initCommand.Flags().StringVarP(&baseURL, "url", "u", baseURL, "define the API base urls")
	initCommand.Flags().StringVarP(&testCasesPath, "test-cases-path", "d", testCasesPath, "define the path to where test scenarios are stored")

	_ = initCommand.MarkFlagRequired("url")
}

func runInitCommand(cmd *cobra.Command, _ []string) error {
	cnf, err := configs.NewConfig(baseURL, testCasesPath)
	if err != nil {
		return err
	}

	confData, err := yaml.Marshal(cnf)
	if err != nil {
		return err
	}

	if err = os.WriteFile(configFile, confData, 0o644); err != nil { //nolint:gosec // we need the permission to be 644
		return err
	}

	cmd.Printf("the config file [%s] successfully created", configFile)

	return nil
}
