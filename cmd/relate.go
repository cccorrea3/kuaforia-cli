package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var relateCmd = &cobra.Command{
	Use:   "relate",
	Short: "Gestionar relaciones entre casos",
}

var relateAddCmd = &cobra.Command{
	Use:   "add [case-id] [related-id]",
	Short: "Crear relación entre dos casos",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		relType, _ := cmd.Flags().GetString("type")
		notes, _ := cmd.Flags().GetString("notes")

		_, err := getClient().Call("create_case_relation", map[string]any{
			"caseId":        parseIntArg(args[0]),
			"relatedCaseId": parseIntArg(args[1]),
			"relationType":  relType,
			"notes":         notes,
		})
		if err != nil {
			return err
		}

		fmt.Printf("Relación creada: %s → %s (%s)\n", args[0], args[1], relType)
		return nil
	},
}

var relateRemoveCmd = &cobra.Command{
	Use:   "remove [case-id] [related-id]",
	Short: "Eliminar relación entre casos",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := getClient().Call("remove_case_relation", map[string]any{
			"caseId":        parseIntArg(args[0]),
			"relatedCaseId": parseIntArg(args[1]),
		})
		if err != nil {
			return err
		}

		fmt.Println("Relación eliminada.")
		return nil
	},
}

func init() {
	relateAddCmd.Flags().String("type", "related", "Tipo: depends_on, contradicts, extends, sucede_a, requiere_datos, related")
	relateAddCmd.Flags().String("notes", "", "Notas sobre la relación")

	relateCmd.AddCommand(relateAddCmd)
	relateCmd.AddCommand(relateRemoveCmd)
	rootCmd.AddCommand(relateCmd)
}
