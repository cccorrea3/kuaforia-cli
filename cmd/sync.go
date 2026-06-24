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

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sincronizar casos entre ./kuaforia/ y el workspace remoto",
	RunE: func(cmd *cobra.Command, args []string) error {
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		pull, _ := cmd.Flags().GetBool("pull")

		if pull {
			return syncPull(dryRun)
		}
		return syncPush(dryRun)
	},
}

var syncWatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Modo watch: sincronización automática al detectar cambios",
	RunE: func(cmd *cobra.Command, args []string) error {
		getFormatter().Warning("Modo watch no disponible: requiere dependencia fsnotify. Usar: while inotifywait -r kuaforia/; do kuaforia sync; done")
		return nil
	},
}

func syncPull(dryRun bool) error {
	f := getFormatter()

	result, err := getClient().Call("export_workspace", map[string]any{
		"workspaceId": getWorkspace(),
		"include":     []string{"cases"},
	})
	if err != nil {
		return fmt.Errorf("error pulling workspace: %w", err)
	}

	var export struct {
		Cases []map[string]any `json:"cases"`
	}
	if err := json.Unmarshal(result, &export); err != nil {
		return fmt.Errorf("error parsing export: %w", err)
	}

	dir := "kuaforia"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating dir: %w", err)
	}

	for _, c := range export.Cases {
		slug := "untitled"
		if s, ok := c["slug"].(string); ok && s != "" {
			slug = s
		} else if t, ok := c["title"].(string); ok && t != "" {
			slug = strings.ReplaceAll(strings.ToLower(t), " ", "-")
		}

		path := filepath.Join(dir, slug+".yaml")

		if dryRun {
			f.Warning("[DRY-RUN] Se escribiría: " + path)
			continue
		}

		out, _ := yaml.Marshal(c)
		if err := os.WriteFile(path, out, 0644); err != nil {
			f.Error("Error escribiendo " + path + ": " + err.Error())
			continue
		}
		f.Success("Escrito: " + path)
	}

	return nil
}

func syncPush(dryRun bool) error {
	dir := "kuaforia"
	info, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("directorio %s no encontrado: %w", dir, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s no es un directorio", dir)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	var cases []map[string]any
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(e.Name()))
		if ext != ".yaml" && ext != ".yml" && ext != ".json" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}

		var caseData map[string]any
		if err := yaml.Unmarshal(data, &caseData); err != nil {
			if err2 := json.Unmarshal(data, &caseData); err2 != nil {
				continue
			}
		}
		cases = append(cases, caseData)
	}

	if len(cases) == 0 {
		return fmt.Errorf("no se encontraron casos en %s", dir)
	}

	params := map[string]any{
		"workspaceId": getWorkspace(),
		"data": map[string]any{
			"cases": cases,
		},
		"mode":   "create_or_update",
		"dryRun": dryRun,
	}

	result, err := getClient().Call("import_workspace", params)
	if err != nil {
		return fmt.Errorf("error syncing: %w", err)
	}

	var stats struct {
		Stats map[string]any `json:"stats"`
	}
	json.Unmarshal(result, &stats)

	if dryRun {
		getFormatter().JSON(stats.Stats)
	} else {
		getFormatter().JSON(stats.Stats)
	}

	return nil
}

func init() {
	syncCmd.Flags().Bool("pull", false, "Bajar cambios remotos → local")
	syncCmd.Flags().Bool("dry-run", false, "Mostrar diff sin aplicar")
	syncCmd.AddCommand(syncWatchCmd)
	rootCmd.AddCommand(syncCmd)
}
