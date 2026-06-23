package cmd

import "github.com/spf13/cobra"

var testsCmd = &cobra.Command{
	Use:   "tests",
	Short: "Gestionar tests de comportamiento",
}

func init() {
	rootCmd.AddCommand(testsCmd)
}
