package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var importCmd = &cobra.Command{
	Use:   "import [file-or-dir]",
	Short: "Importar casos desde archivos YAML/JSON",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dryRun, _ := cmd.Flags().GetBool("dry-run")
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
				if ext == ".json" || ext == ".yaml" || ext == ".yml" {
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

		for _, file := range files {
			data, err := os.ReadFile(file)
			if err != nil {
				f.Error("Error leyendo " + filepath.Base(file) + ": " + err.Error())
				continue
			}

			var caseData map[string]any
			if err := yaml.Unmarshal(data, &caseData); err != nil {
				if err2 := json.Unmarshal(data, &caseData); err2 != nil {
					f.Error("Error parseando " + filepath.Base(file) + " (YAML/JSON)")
					continue
				}
			}

			caseData["workspaceId"] = getWorkspace()

			title := caseData["title"]
			if title == nil {
				title = filepath.Base(file)
			}

			if dryRun {
				f.Warning("[DRY-RUN] Se crearía: " + fmt.Sprint(title))
				continue
			}

			result, err := getClient().Call("create_case", caseData)
			if err != nil {
				f.Error("Error creando '" + fmt.Sprint(title) + "': " + err.Error())
				continue
			}

			f.Success("Creado desde: " + filepath.Base(file))
			_ = result
		}

		return nil
	},
}

func init() {
	importCmd.Flags().Bool("dry-run", false, "Mostrar qué se crearía sin ejecutar")
	rootCmd.AddCommand(importCmd)
}
