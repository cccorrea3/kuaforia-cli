package cmd

import (
	"github.com/spf13/cobra"
)

var testsRunCmd = &cobra.Command{
	Use:   "run [case-id]",
	Short: "Ejecutar tests de comportamiento",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, _ := cmd.Flags().GetBool("all")

		if all {
			result, err := getClient().Call("run_all_case_tests", map[string]any{})
			if err != nil {
				return err
			}
			getFormatter().JSON(result)
			return nil
		}

		if len(args) < 1 {
			return cmd.Help()
		}

		params := map[string]any{"caseId": parseIntArg(args[0])}

		if cmd.Flag("test").Changed {
			testID, _ := cmd.Flags().GetInt("test")
			params["testId"] = testID
			result, err := getClient().Call("run_case_test", params)
			if err != nil {
				return err
			}
			getFormatter().JSON(result)
			return nil
		}

		result, err := getClient().Call("run_all_case_tests", params)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	testsRunCmd.Flags().Int("test", 0, "ID del test específico a ejecutar")
	testsRunCmd.Flags().Bool("all", false, "Ejecutar todos los tests del workspace")
	testsCmd.AddCommand(testsRunCmd)
}
