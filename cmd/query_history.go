package cmd

import (
	"github.com/spf13/cobra"
)

var queryHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Ver historial de consultas",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]any{"workspaceId": getWorkspace()}

		if cmd.Flag("limit").Changed {
			params["limit"] = parseIntFlag(cmd, "limit")
		}
		if cmd.Flag("failed").Changed {
			params["failed"] = true
		}

		result, err := getClient().Call("get_query_history", params)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	queryHistoryCmd.Flags().Int("limit", 20, "Máximo de resultados")
	queryHistoryCmd.Flags().Bool("failed", false, "Solo consultas fallidas")
	queryCmd.AddCommand(queryHistoryCmd)
}
