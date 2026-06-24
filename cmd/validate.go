package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kuaforia/cli/internal/schema"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var validateCmd = &cobra.Command{
	Use:   "validate [file-or-dir]",
	Short: "Validar archivos .kuaforia.yaml contra el schema",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		strict, _ := cmd.Flags().GetBool("strict")
		path := args[0]

		info, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("error accessing path: %w", err)
		}

		var files []string
		if info.IsDir() {
			entries, err := os.ReadDir(path)
			if err != nil {
				return fmt.Errorf("error reading directory: %w", err)
			}
			for _, e := range entries {
				if e.IsDir() {
					continue
				}
				ext := strings.ToLower(filepath.Ext(e.Name()))
				if ext == ".yaml" || ext == ".yml" || ext == ".json" {
					files = append(files, filepath.Join(path, e.Name()))
				}
			}
		} else {
			files = append(files, path)
		}

		if len(files) == 0 {
			return fmt.Errorf("no se encontraron archivos YAML/JSON en: %s", path)
		}

		f := getFormatter()
		hasErrors := false

		for _, file := range files {
			data, err := os.ReadFile(file)
			if err != nil {
				f.Error("Error leyendo " + filepath.Base(file) + ": " + err.Error())
				hasErrors = true
				continue
			}

			var doc map[string]any
			if err := yaml.Unmarshal(data, &doc); err != nil {
				if err2 := json.Unmarshal(data, &doc); err2 != nil {
					f.Error("Error parseando " + filepath.Base(file) + " (YAML/JSON)")
					hasErrors = true
					continue
				}
			}

			errs := schema.Validate(doc)
			if len(errs) > 0 {
				hasErrors = true
				for _, e := range errs {
					f.Error(filepath.Base(file) + ": " + e.Field + " — " + e.Message)
				}
			} else {
				f.Success("OK: " + filepath.Base(file))
			}
		}

		if hasErrors && strict {
			return fmt.Errorf("validación falló (modo strict)")
		}
		return nil
	},
}

func init() {
	validateCmd.Flags().Bool("strict", false, "Falla con código de error si hay warnings")
	rootCmd.AddCommand(validateCmd)
}
