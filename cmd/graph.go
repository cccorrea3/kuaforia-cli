package cmd

import (
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Visualizar grafo de dependencias",
}

var graphDepsCmd = &cobra.Command{
	Use:   "dependencies [case-id]",
	Short: "Mostrar dependencias de un caso como árbol",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		depth, _ := cmd.Flags().GetInt("depth")

		result, err := getClient().Call("get_dependency_graph", map[string]any{
			"caseId": parseIntArg(args[0]),
			"depth":  depth,
		})
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

var graphCriticalCmd = &cobra.Command{
	Use:   "critical [case-id]",
	Short: "Mostrar la ruta crítica del grafo",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("get_critical_path", map[string]any{
			"caseId": parseIntArg(args[0]),
		})
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

var graphSpofCmd = &cobra.Command{
	Use:   "spof",
	Short: "Detectar single points of failure en el workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("get_graph_metrics", map[string]any{
			"workspaceId": getWorkspace(),
		})
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	graphDepsCmd.Flags().Int("depth", 3, "Profundidad máxima del árbol")

	graphCmd.AddCommand(graphDepsCmd)
	graphCmd.AddCommand(graphCriticalCmd)
	graphCmd.AddCommand(graphSpofCmd)
	rootCmd.AddCommand(graphCmd)
}
