package cmd

import (
	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Gestionar versiones de casos",
}

var versionsListCmd = &cobra.Command{
	Use:   "list [case-id]",
	Short: "Listar versiones de un caso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("get_case_versions", map[string]any{
			"caseId": parseIntArg(args[0]),
		})
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

var versionsDiffCmd = &cobra.Command{
	Use:   "diff [case-id]",
	Short: "Mostrar diferencias entre versiones",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]any{"caseId": parseIntArg(args[0])}

		if cmd.Flag("from").Changed {
			params["fromVersion"] = parseIntFlag(cmd, "from")
		}
		if cmd.Flag("to").Changed {
			params["toVersion"] = parseIntFlag(cmd, "to")
		}

		result, err := getClient().Call("get_case_diff", params)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	versionsDiffCmd.Flags().Int("from", 0, "Versión origen (default: última)")
	versionsDiffCmd.Flags().Int("to", 0, "Versión destino (default: actual)")

	versionsCmd.AddCommand(versionsListCmd)
	versionsCmd.AddCommand(versionsDiffCmd)
	rootCmd.AddCommand(versionsCmd)
}
