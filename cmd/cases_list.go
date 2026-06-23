package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

type caseRow struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
	Category string `json:"category"`
}

var casesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar casos de negocio",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := getClient()
		result, err := c.Call("list_cases", map[string]any{
			"workspaceId": getWorkspace(),
			"status":      cmd.Flag("status").Value.String(),
			"priority":    cmd.Flag("priority").Value.String(),
			"search":      cmd.Flag("search").Value.String(),
			"page":        1,
			"perPage":     100,
		})
		if err != nil {
			return err
		}

		var cases []caseRow
		if err := json.Unmarshal(result, &cases); err != nil {
			var wrapper struct {
				Data []caseRow `json:"data"`
			}
			if err2 := json.Unmarshal(result, &wrapper); err2 != nil {
				return fmt.Errorf("error parsing response: %w", err)
			}
			cases = wrapper.Data
		}

		f := getFormatter()
		f.Table(
			[]string{"ID", "TITLE", "STATUS", "PRIORITY", "CATEGORY"},
			toRows(cases),
		)

		fmt.Printf("Total: %d casos\n", len(cases))
		return nil
	},
}

func init() {
	casesListCmd.Flags().String("status", "", "Filtrar por estado")
	casesListCmd.Flags().String("priority", "", "Filtrar por prioridad")
	casesListCmd.Flags().String("search", "", "Búsqueda por texto")
	casesListCmd.Flags().Int("page", 1, "Número de página")
	casesListCmd.Flags().Int("per-page", 20, "Resultados por página")
	casesCmd.AddCommand(casesListCmd)
}

func toRows(cases []caseRow) [][]string {
	rows := make([][]string, len(cases))
	for i, c := range cases {
		if c.Category == "" {
			c.Category = "-"
		}
		rows[i] = []string{
			strconv.Itoa(c.ID),
			c.Title,
			c.Status,
			c.Priority,
			c.Category,
		}
	}
	return rows
}
