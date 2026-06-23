package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var casesDeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Eliminar lógicamente un caso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")
		caseID := parseIntArg(args[0])

		if !force {
			result, err := getClient().Call("get_case", map[string]any{"caseId": caseID})
			if err == nil {
				f := getFormatter()
				f.JSON(result)
			}
			fmt.Println()

			if !promptConfirm("Eliminar este caso?") {
				fmt.Println("Cancelado.")
				return nil
			}
		}

		_, err := getClient().Call("delete_case", map[string]any{"caseId": caseID})
		if err != nil {
			return err
		}

		fmt.Printf("Caso %d eliminado.\n", caseID)
		return nil
	},
}

func init() {
	casesDeleteCmd.Flags().Bool("force", false, "Omitir confirmación")
	casesCmd.AddCommand(casesDeleteCmd)
}

func promptConfirm(label string) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	_, err := prompt.Run()
	return err == nil
}
