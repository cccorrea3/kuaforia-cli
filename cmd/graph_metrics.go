package cmd

import (
	"github.com/spf13/cobra"
)

var graphMetricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Ver métricas del grafo de dependencias",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := map[string]any{"workspaceId": getWorkspace()}

		if cmd.Flag("sort").Changed {
			params["sort"] = cmd.Flag("sort").Value.String()
		}
		if cmd.Flag("desc").Changed {
			params["desc"] = true
		}
		if cmd.Flag("filter").Changed {
			params["filter"] = cmd.Flag("filter").Value.String()
		}

		result, err := getClient().Call("get_graph_metrics", params)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	graphMetricsCmd.Flags().String("sort", "", "Ordenar por: betweenness, degree, closeness")
	graphMetricsCmd.Flags().Bool("desc", false, "Orden descendente")
	graphMetricsCmd.Flags().String("filter", "", "Filtrar: spof, isolated")
	graphCmd.AddCommand(graphMetricsCmd)
}
