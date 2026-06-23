package cmd

import (
	"github.com/spf13/cobra"
)

var impactCmd = &cobra.Command{
	Use:   "impact",
	Short: "Analizar impacto de cambios en casos",
}

var impactAnalyzeCmd = &cobra.Command{
	Use:   "analyze [case-id]",
	Short: "Simular impacto de eliminar/desactivar un caso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		scenario, _ := cmd.Flags().GetString("scenario")

		result, err := getClient().Call("analyze_impact", map[string]any{
			"caseId":   parseIntArg(args[0]),
			"scenario": scenario,
		})
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

var impactQuickCmd = &cobra.Command{
	Use:   "quick [case-id]",
	Short: "Análisis rápido de impacto",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("quick_impact_analysis", map[string]any{
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
	impactAnalyzeCmd.Flags().String("scenario", "deletion", "Escenario: deletion, deactivation")

	impactCmd.AddCommand(impactAnalyzeCmd)
	impactCmd.AddCommand(impactQuickCmd)
	rootCmd.AddCommand(impactCmd)
}
