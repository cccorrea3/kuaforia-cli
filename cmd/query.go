package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query [question]",
	Short: "Ejecutar una consulta al motor de IA",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("run_query", map[string]any{
			"workspaceId": getWorkspace(),
			"question":    strings.Join(args, " "),
		})
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
