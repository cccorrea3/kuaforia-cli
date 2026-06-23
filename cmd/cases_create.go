package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var casesCreateCmd = &cobra.Command{
	Use:   "create [title]",
	Short: "Crear un nuevo caso de negocio",
	Long: `Crea un caso en estado draft. Los campos pueden pasarse como flags
o mediante un archivo YAML/JSON con --file.

Ejemplos:
  kuaforia cases create "Error en login" --description "..." --priority high
  kuaforia cases create --file caso.yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var caseData map[string]any

		if filePath := cmd.Flag("file").Value.String(); filePath != "" {
			data, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("error reading file: %w", err)
			}
			if err := yaml.Unmarshal(data, &caseData); err != nil {
				if err2 := json.Unmarshal(data, &caseData); err2 != nil {
					return fmt.Errorf("error parsing file (YAML/JSON): %w", err)
				}
			}
		} else {
			if len(args) < 1 {
				return fmt.Errorf("title is required when not using --file")
			}
			caseData = map[string]any{
				"title":             args[0],
				"description":       cmd.Flag("description").Value.String(),
				"caseTypeId":        parseIntFlag(cmd, "type-id"),
				"categoryId":        parseIntFlag(cmd, "category-id"),
				"triggerContext":    cmd.Flag("trigger-context").Value.String(),
				"expectedBehavior":  cmd.Flag("expected-behavior").Value.String(),
				"tags":              splitFlag(cmd, "tags"),
				"priority":          cmd.Flag("priority").Value.String(),
				"systemModule":      cmd.Flag("module").Value.String(),
				"systemVersion":     cmd.Flag("version").Value.String(),
			}
		}

		caseData["workspaceId"] = getWorkspace()

		result, err := getClient().Call("create_case", caseData)
		if err != nil {
			return err
		}

		f := getFormatter()
		f.JSON(result)
		return nil
	},
}

func init() {
	casesCreateCmd.Flags().String("description", "", "Descripción detallada")
	casesCreateCmd.Flags().Int("type-id", 0, "ID del tipo de caso")
	casesCreateCmd.Flags().Int("category-id", 0, "ID de la categoría")
	casesCreateCmd.Flags().String("trigger-context", "", "Contexto que dispara el caso")
	casesCreateCmd.Flags().String("expected-behavior", "", "Comportamiento esperado")
	casesCreateCmd.Flags().String("tags", "", "Tags separados por coma")
	casesCreateCmd.Flags().String("priority", "medium", "Prioridad: low, medium, high, critical")
	casesCreateCmd.Flags().String("module", "", "Módulo del sistema")
	casesCreateCmd.Flags().String("version", "", "Versión del sistema")
	casesCreateCmd.Flags().String("file", "", "Archivo YAML/JSON con datos del caso")
	casesCmd.AddCommand(casesCreateCmd)
}

func parseIntFlag(cmd *cobra.Command, name string) int {
	v, _ := strconv.Atoi(cmd.Flag(name).Value.String())
	return v
}
