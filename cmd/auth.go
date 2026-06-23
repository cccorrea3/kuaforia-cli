package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kuaforia/cli/internal/client"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Gestionar autenticación",
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Iniciar sesión y guardar API key",
	RunE: func(cmd *cobra.Command, args []string) error {
		prompt := promptui.Prompt{Label: "Server URL", Default: viper.GetString("server")}
		server, err := prompt.Run()
		if err != nil {
			return err
		}

		prompt = promptui.Prompt{Label: "Tenant slug", Default: viper.GetString("tenant")}
		tenant, err := prompt.Run()
		if err != nil {
			return err
		}

		prompt = promptui.Prompt{Label: "API key", Mask: '*', Default: viper.GetString("api_key")}
		apiKey, err := prompt.Run()
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Print("Validando conexión... ")

		c := client.New(server, apiKey, tenant)
		health, err := c.Get("health")
		if err != nil {
			fmt.Println("[ERROR]")
			return fmt.Errorf("no se pudo conectar al servidor: %w", err)
		}

		var data map[string]any
		json.Unmarshal(health, &data)
		fmt.Println("[OK]")
		fmt.Printf("  Key name: %v\n", data["key_name"])
		fmt.Printf("  Tenant:   %v\n", data["tenant"])
		fmt.Printf("  Scopes:   %v\n", data["scopes"])

		if err := saveConfig(server, tenant, apiKey); err != nil {
			return fmt.Errorf("no se pudo guardar la config: %w", err)
		}

		fmt.Println()
		fmt.Println("Configuración guardada correctamente.")
		return nil
	},
}

var authSetKeyCmd = &cobra.Command{
	Use:   "set-key [key]",
	Short: "Establecer API key en la configuración",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		server := viper.GetString("server")
		tenant := viper.GetString("tenant")

		if server == "" {
			server = "http://localhost:8000"
		}

		if err := saveConfig(server, tenant, args[0]); err != nil {
			return fmt.Errorf("no se pudo guardar la config: %w", err)
		}

		fmt.Println("API key guardada correctamente.")
		return nil
	},
}

var authTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Probar conexión con el servidor",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := viper.GetString("server")
		tenant := viper.GetString("tenant")
		apiKey := viper.GetString("api_key")

		if server == "" || apiKey == "" {
			return fmt.Errorf("ejecuta 'kuaforia auth login' para configurar la conexión")
		}

		fmt.Printf("Server: %s\n", server)
		fmt.Printf("Tenant: %s\n", tenant)
		fmt.Printf("API Key: %s…\n\n", safePrefix(apiKey, 12))

		fmt.Print("Verificando conexión... ")

		c := client.New(server, apiKey, tenant)
		health, err := c.Get("health")
		if err != nil {
			fmt.Println("[ERROR]")
			return fmt.Errorf("error de conexión: %w", err)
		}

		fmt.Println("[OK]")
		fmt.Println()

		var data map[string]any
		json.Unmarshal(health, &data)

		fmt.Printf("  Key name: %v\n", data["key_name"])
		fmt.Printf("  Tenant:   %v\n", data["tenant"])
		fmt.Printf("  Scopes:   %v\n", data["scopes"])

		fmt.Println()
		fmt.Println("Conexión exitosa.")
		return nil
	},
}

func init() {
	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authSetKeyCmd)
	authCmd.AddCommand(authTestCmd)
	rootCmd.AddCommand(authCmd)
}

type configFile struct {
	Server           string `yaml:"server" json:"server"`
	Tenant           string `yaml:"tenant" json:"tenant"`
	APIKey           string `yaml:"api_key" json:"api_key"`
	Output           string `yaml:"output" json:"output"`
	DefaultWorkspace int    `yaml:"default_workspace" json:"default_workspace"`
}

func saveConfig(server, tenant, apiKey string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, ".kuaforia.yaml")

	cfg := configFile{
		Server:           server,
		Tenant:           tenant,
		APIKey:           apiKey,
		Output:           viper.GetString("output"),
		DefaultWorkspace: viper.GetInt("default_workspace"),
	}

	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

func safePrefix(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n]
}


