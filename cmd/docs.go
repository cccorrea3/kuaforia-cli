package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generar documentación markdown de comandos",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, _ := cmd.Flags().GetString("dir")
		if dir == "" {
			return fmt.Errorf("flag --dir es requerido")
		}
		return doc.GenMarkdownTree(rootCmd, dir)
	},
}

func init() {
	docsCmd.Flags().String("dir", "./docs/cli/", "Directorio de salida para la documentación")
	rootCmd.AddCommand(docsCmd)
}
