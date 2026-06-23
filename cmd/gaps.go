package cmd

import (
	"github.com/spf13/cobra"
)

var gapsCmd = &cobra.Command{
	Use:   "gaps",
	Short: "Gestionar knowledge gaps",
}

var gapsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar knowledge gaps detectados",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]any{"workspaceId": getWorkspace()}

		if cmd.Flag("priority").Changed {
			params["priority"] = cmd.Flag("priority").Value.String()
		}
		if cmd.Flag("status").Changed {
			params["status"] = cmd.Flag("status").Value.String()
		}

		result, err := getClient().Call("get_knowledge_gaps", params)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

var gapsMarkCmd = &cobra.Command{
	Use:   "mark [id]",
	Short: "Marcar gap como completado o descartado",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		getFormatter().Warning("Comando no implementado: no existe tool MCP para marcar gaps")
		return nil
	},
}

func init() {
	gapsListCmd.Flags().String("priority", "", "Filtrar por prioridad: high, medium, low")
	gapsListCmd.Flags().String("status", "", "Filtrar por estado: pending, completed, dismissed")

	gapsMarkCmd.Flags().Bool("completed", false, "Marcar como completado")
	gapsMarkCmd.Flags().Bool("dismissed", false, "Descartar gap")

	gapsCmd.AddCommand(gapsListCmd)
	gapsCmd.AddCommand(gapsMarkCmd)
	rootCmd.AddCommand(gapsCmd)
}
