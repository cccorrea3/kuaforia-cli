package cmd

import (
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exportar datos del workspace",
}

var exportCasesCmd = &cobra.Command{
	Use:   "cases",
	Short: "Exportar casos del workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]any{"workspaceId": getWorkspace()}

		include := []string{"cases", "relations", "tests", "versions", "categories", "types", "custom_fields"}
		if cmd.Flag("include").Changed {
			include = splitFlag(cmd, "include")
		}
		params["include"] = include

		result, err := getClient().Call("export_workspace", params)
		if err != nil {
			return err
		}

		if cmd.Flag("format").Changed && cmd.Flag("format").Value.String() == "json" {
			getFormatter().JSON(result)
		} else {
			getFormatter().JSON(result)
		}
		return nil
	},
}

var exportReportCmd = &cobra.Command{
	Use:   "report [type]",
	Short: "Generar reporte del workspace",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		getFormatter().Warning("Reportes no disponibles: no existe tool MCP de exportación de reportes")
		return nil
	},
}

func init() {
	exportCasesCmd.Flags().StringSlice("include", nil, "Qué incluir: cases,relations,tests,versions,categories,types,custom_fields")
	exportCasesCmd.Flags().String("format", "json", "Formato: json")

	exportReportCmd.Flags().String("format", "html", "Formato: html, pdf")

	exportCmd.AddCommand(exportCasesCmd)
	exportCmd.AddCommand(exportReportCmd)
	rootCmd.AddCommand(exportCmd)
}
