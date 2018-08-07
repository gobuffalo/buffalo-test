package cmd

import (
	"os"

	"github.com/gobuffalo/buffalo-test/test"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var forceMigrations = false

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:                "test",
	Short:              "Runs the tests for your Buffalo app",
	DisableFlagParsing: true,
	RunE:               runTests,
}

func runTests(cmd *cobra.Command, args []string) error {
	logrus.Info("Running plugin version of test command")
	os.Setenv("GO_ENV", "test")

	if _, err := os.Stat("database.yml"); err != nil {
		return test.NewRunner(args).Run()
	}

	sup, err := test.NewSetup(args)
	if err != nil {
		return err
	}

	if err := sup.Run(); err != nil {
		return err
	}

	args = removeFlag("--force-migrations", args)
	args = removeFlag("--skip-db-setup", args)

	return test.NewRunner(args).Run()
}

func removeFlag(flag string, args []string) []string {
	for i, v := range args {
		if v != flag {
			continue
		}

		args = append(args[:i], args[i+1:]...)
		break
	}

	return args
}

func init() {
	rootCmd.AddCommand(testCmd)
}
