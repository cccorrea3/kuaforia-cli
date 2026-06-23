package cmd

import "github.com/spf13/cobra"

var casesCmd = &cobra.Command{
	Use:   "cases",
	Short: "Gestionar casos de negocio",
}

func init() {
	rootCmd.AddCommand(casesCmd)
}
