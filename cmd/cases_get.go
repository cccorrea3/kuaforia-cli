package cmd

import (
	"github.com/spf13/cobra"
)

var casesGetCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Obtener detalle completo de un caso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("get_case", map[string]any{
			"caseId": parseIntArg(args[0]),
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	casesCmd.AddCommand(casesGetCmd)
}
