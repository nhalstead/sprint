package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sprint",
		Long: "Get a certificate fingerprint from a FQDN or URL.",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(fingerprintCmd)
	rootCmd.AddCommand(matchCmd)
}
