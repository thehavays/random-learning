package cmd

import (
	"github.com/spf13/cobra"
)

// Root command
var rootCmd = &cobra.Command{
	Use:   "random-learning",
	Short: "A CLI tool for interacting with StackExchange",
}

// Execute the root command
func Execute() error {
	return rootCmd.Execute()
}
