package cmd

import (
	"github.com/spf13/cobra"
)

var casesDuplicateCmd = &cobra.Command{
	Use:   "duplicate [id]",
	Short: "Duplicar un caso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]any{"caseId": parseIntArg(args[0])}

		if cmd.Flag("new-title").Changed {
			params["newTitle"] = cmd.Flag("new-title").Value.String()
		}

		result, err := getClient().Call("duplicate_case", params)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	casesDuplicateCmd.Flags().String("new-title", "", "Título del caso duplicado")
	casesCmd.AddCommand(casesDuplicateCmd)
}
