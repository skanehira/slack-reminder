package cmd

import (
	"fmt"
	"os"

	"github.com/skanehira/slack-reminder/survey"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "slack-reminder",
}

func exitError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func Execute() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if err := survey.Ask(); err != nil {
			exitError(err)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}
