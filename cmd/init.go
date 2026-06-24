package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//go:embed templates/kuaforia.yaml
var kuaforiaYaml string

//go:embed templates/readme.md
var readmeMd string

//go:embed templates/ejemplo-basico.yaml
var ejemploBasicoYaml string

//go:embed templates/ejemplo-completo.yaml
var ejemploCompletoYaml string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializar estructura ./kuaforia/ para casos como código",
	RunE: func(cmd *cobra.Command, args []string) error {
		full, _ := cmd.Flags().GetBool("template")

		dirs := []string{"kuaforia/examples", "kuaforia/casos"}
		for _, d := range dirs {
			if err := os.MkdirAll(d, 0755); err != nil {
				return fmt.Errorf("error creating %s: %w", d, err)
			}
		}

		files := map[string]string{
			"kuaforia/.kuaforia.yaml":       kuaforiaYaml,
			"kuaforia/README.md":            readmeMd,
			"kuaforia/examples/ejemplo.yaml": ejemploBasicoYaml,
		}

		if full {
			files["kuaforia/examples/ejemplo-completo.yaml"] = ejemploCompletoYaml
		}

		f := getFormatter()

		for path, content := range files {
			if err := os.WriteFile(path, []byte(content), 0644); err != nil {
				f.Error("Error escribiendo " + path + ": " + err.Error())
				continue
			}
			f.Success("Creado: " + path)
		}

		return nil
	},
}

func init() {
	initCmd.Flags().Bool("template", false, "Incluir ejemplos completos")
	rootCmd.AddCommand(initCmd)
}
