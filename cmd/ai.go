package cmd

import (
	"github.com/spf13/cobra"
)

var aiCmd = &cobra.Command{
	Use:   "ai",
	Short: "Gestionar configuración de IA",
}

var aiSettingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Ver configuración actual de IA",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("get_ai_settings", map[string]any{
			"workspaceId": getWorkspace(),
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

var aiConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Cambiar provider/modelo de IA",
	RunE: func(cmd *cobra.Command, args []string) error {
		getFormatter().Warning("Comando no implementado: no existe tool MCP para configurar IA")
		return nil
	},
}

var aiEmbeddingsCmd = &cobra.Command{
	Use:   "embeddings",
	Short: "Gestionar embeddings",
}

var aiEmbeddingsGenerateCmd = &cobra.Command{
	Use:   "generate [case-id]",
	Short: "Regenerar embeddings de uno o todos los casos",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		getFormatter().Warning("Comando no implementado: no existe tool MCP para regenerar embeddings")
		return nil
	},
}

func init() {
	aiConfigureCmd.Flags().String("provider", "", "Provider: ollama, openai, anthropic")
	aiConfigureCmd.Flags().String("model", "", "Modelo (ej: llama3, gpt-4)")

	aiEmbeddingsGenerateCmd.Flags().Bool("all", false, "Regenerar todos los embeddings")

	aiEmbeddingsCmd.AddCommand(aiEmbeddingsGenerateCmd)
	aiCmd.AddCommand(aiSettingsCmd)
	aiCmd.AddCommand(aiConfigureCmd)
	aiCmd.AddCommand(aiEmbeddingsCmd)
	rootCmd.AddCommand(aiCmd)
}
