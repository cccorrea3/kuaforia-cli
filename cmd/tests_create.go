package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var testsCreateCmd = &cobra.Command{
	Use:   "create [case-id]",
	Short: "Crear un test de comportamiento para un caso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var testData map[string]any

		if filePath := cmd.Flag("file").Value.String(); filePath != "" {
			data, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("error reading file: %w", err)
			}
			if err := yaml.Unmarshal(data, &testData); err != nil {
				if err2 := json.Unmarshal(data, &testData); err2 != nil {
					return fmt.Errorf("error parsing file (YAML/JSON): %w", err)
				}
			}
		} else {
			testData = map[string]any{
				"caseId":          parseIntArg(args[0]),
				"type":            cmd.Flag("type").Value.String(),
				"question":        cmd.Flag("question").Value.String(),
				"expectedTerms":   splitFlag(cmd, "expected-terms"),
				"minConfidence":   0.7,
			}
			if cmd.Flag("min-confidence").Changed {
				testData["minConfidence"] = parseFloatFlag(cmd, "min-confidence")
			}
		}

		result, err := getClient().Call("create_case_test", testData)
		if err != nil {
			return err
		}

		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	testsCreateCmd.Flags().String("type", "query_result", "Tipo de test: query_result, ...")
	testsCreateCmd.Flags().String("question", "", "Pregunta del test")
	testsCreateCmd.Flags().String("expected-terms", "", "Términos esperados separados por coma")
	testsCreateCmd.Flags().Float64("min-confidence", 0.7, "Confianza mínima (0-1)")
	testsCreateCmd.Flags().String("file", "", "Archivo YAML/JSON con definición del test")
	testsCmd.AddCommand(testsCreateCmd)
}
