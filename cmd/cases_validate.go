package cmd

import (
	"github.com/spf13/cobra"
)

var casesValidateCmd = &cobra.Command{
	Use:   "validate [id]",
	Short: "Validar o rechazar un caso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]any{"caseId": parseIntArg(args[0])}

		if cmd.Flag("status").Changed {
			params["status"] = cmd.Flag("status").Value.String()
		}
		if cmd.Flag("notes").Changed {
			params["notes"] = cmd.Flag("notes").Value.String()
		}

		result, err := getClient().Call("validate_case", params)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	casesValidateCmd.Flags().String("status", "", "approved, rejected, or needs_changes")
	casesValidateCmd.Flags().String("notes", "", "Notas de validación")
	casesCmd.AddCommand(casesValidateCmd)
}
