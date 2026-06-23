package cmd

import (
	"github.com/spf13/cobra"
)

var casesPublishCmd = &cobra.Command{
	Use:   "publish [id]",
	Short: "Publicar una nueva versión de un caso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]any{"caseId": parseIntArg(args[0])}

		if cmd.Flag("changelog").Changed {
			params["changelog"] = cmd.Flag("changelog").Value.String()
		}
		if cmd.Flag("deprecation").Changed {
			params["deprecation"] = cmd.Flag("deprecation").Value.String()
		}

		result, err := getClient().Call("publish_case", params)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	casesPublishCmd.Flags().String("changelog", "", "Descripción de cambios")
	casesPublishCmd.Flags().String("deprecation", "", "Mensaje de deprecación (opcional)")
	casesCmd.AddCommand(casesPublishCmd)
}
