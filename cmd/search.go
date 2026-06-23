package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Buscar casos por texto",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]any{
			"workspaceId": getWorkspace(),
			"query":       strings.Join(args, " "),
		}

		if cmd.Flag("limit").Changed {
			params["limit"] = parseIntFlag(cmd, "limit")
		}
		if cmd.Flag("category").Changed {
			params["category"] = cmd.Flag("category").Value.String()
		}

		result, err := getClient().Call("search_cases", params)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	searchCmd.Flags().Int("limit", 20, "Máximo de resultados")
	searchCmd.Flags().String("category", "", "Filtrar por categoría")
	rootCmd.AddCommand(searchCmd)
}
