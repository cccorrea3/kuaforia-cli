package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var typesCmd = &cobra.Command{
	Use:   "types",
	Short: "Gestionar tipos de caso",
}

var typesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar tipos de caso disponibles",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("list_case_types", map[string]any{
			"workspaceId": getWorkspace(),
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Gestionar categorías",
}

var categoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar categorías disponibles",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("list_case_categories", map[string]any{
			"workspaceId": getWorkspace(),
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Gestionar plantillas",
}

var templatesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar plantillas disponibles",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("list_templates", map[string]any{
			"workspaceId": getWorkspace(),
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

var templatesApplyCmd = &cobra.Command{
	Use:   "apply [case-id]",
	Short: "Aplicar plantilla a un caso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		templateID, _ := cmd.Flags().GetInt("template")
		if templateID == 0 {
			return fmt.Errorf("flag --template es requerido")
		}

		result, err := getClient().Call("apply_template", map[string]any{
			"workspaceId": getWorkspace(),
			"caseId":      parseIntArg(args[0]),
			"templateId":  templateID,
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

var templatesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Crear nueva plantilla desde archivo YAML/JSON",
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath, _ := cmd.Flags().GetString("file")
		if filePath == "" {
			return fmt.Errorf("flag --file es requerido")
		}

		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		var templateData map[string]any
		if err := yaml.Unmarshal(data, &templateData); err != nil {
			if err2 := json.Unmarshal(data, &templateData); err2 != nil {
				return fmt.Errorf("error parsing file (YAML/JSON): %w", err)
			}
		}

		templateData["workspaceId"] = getWorkspace()

		result, err := getClient().Call("create_template", templateData)
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	templatesApplyCmd.Flags().Int("template", 0, "ID de la plantilla a aplicar")
	templatesCreateCmd.Flags().String("file", "", "Archivo YAML/JSON con la plantilla")

	typesCmd.AddCommand(typesListCmd)
	categoriesCmd.AddCommand(categoriesListCmd)
	templatesCmd.AddCommand(templatesListCmd)
	templatesCmd.AddCommand(templatesApplyCmd)
	templatesCmd.AddCommand(templatesCreateCmd)

	rootCmd.AddCommand(typesCmd)
	rootCmd.AddCommand(categoriesCmd)
	rootCmd.AddCommand(templatesCmd)
}
