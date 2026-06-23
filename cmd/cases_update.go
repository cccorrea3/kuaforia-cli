package cmd

import (
	"github.com/spf13/cobra"
)

var casesUpdateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Actualizar campos de un caso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		updates := map[string]any{"caseId": parseIntArg(args[0])}

		if cmd.Flag("title").Changed {
			updates["title"] = cmd.Flag("title").Value.String()
		}
		if cmd.Flag("description").Changed {
			updates["description"] = cmd.Flag("description").Value.String()
		}
		if cmd.Flag("status").Changed {
			updates["status"] = cmd.Flag("status").Value.String()
		}
		if cmd.Flag("priority").Changed {
			updates["priority"] = cmd.Flag("priority").Value.String()
		}
		if cmd.Flag("tags").Changed {
			updates["tags"] = splitFlag(cmd, "tags")
		}
		if cmd.Flag("type-id").Changed {
			updates["caseTypeId"] = parseIntFlag(cmd, "type-id")
		}
		if cmd.Flag("category-id").Changed {
			updates["categoryId"] = parseIntFlag(cmd, "category-id")
		}
		if cmd.Flag("trigger-context").Changed {
			updates["triggerContext"] = cmd.Flag("trigger-context").Value.String()
		}
		if cmd.Flag("expected-behavior").Changed {
			updates["expectedBehavior"] = cmd.Flag("expected-behavior").Value.String()
		}
		if cmd.Flag("module").Changed {
			updates["systemModule"] = cmd.Flag("module").Value.String()
		}
		if cmd.Flag("version").Changed {
			updates["systemVersion"] = cmd.Flag("version").Value.String()
		}

		result, err := getClient().Call("update_case", updates)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	casesUpdateCmd.Flags().String("title", "", "Nuevo título")
	casesUpdateCmd.Flags().String("description", "", "Nueva descripción")
	casesUpdateCmd.Flags().String("status", "", "Nuevo estado")
	casesUpdateCmd.Flags().String("priority", "", "Nueva prioridad")
	casesUpdateCmd.Flags().String("tags", "", "Tags separados por coma")
	casesUpdateCmd.Flags().Int("type-id", 0, "ID del tipo de caso")
	casesUpdateCmd.Flags().Int("category-id", 0, "ID de la categoría")
	casesUpdateCmd.Flags().String("trigger-context", "", "Contexto que dispara el caso")
	casesUpdateCmd.Flags().String("expected-behavior", "", "Comportamiento esperado")
	casesUpdateCmd.Flags().String("module", "", "Módulo del sistema")
	casesUpdateCmd.Flags().String("version", "", "Versión del sistema")
	casesCmd.AddCommand(casesUpdateCmd)
}
