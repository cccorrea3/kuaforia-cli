package cmd

import (
	"github.com/spf13/cobra"
)

var casesStatusCmd = &cobra.Command{
	Use:   "status [id] [new-status]",
	Short: "Cambiar estado de un caso",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("transition_case_status", map[string]any{
			"caseId": parseIntArg(args[0]),
			"status": args[1],
		})
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	casesCmd.AddCommand(casesStatusCmd)
}
