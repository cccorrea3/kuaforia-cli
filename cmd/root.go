package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "kuaforia",
	Short: "CLI para gestionar Kuaforia",
	Long: `kuaforia — Interfaz de línea de comandos para gestionar
casos de negocio, workspaces, relaciones, pruebas y más
en la plataforma Kuaforia.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default ~/.kuaforia.yaml)")
	rootCmd.PersistentFlags().StringP("tenant", "t", "", "Tenant slug")
	rootCmd.PersistentFlags().StringP("api-key", "k", "", "API key")
	rootCmd.PersistentFlags().StringP("server", "s", "http://localhost:8000", "Server URL")
	rootCmd.PersistentFlags().StringP("output", "o", "table", "Output format: table, json, yaml")

	viper.BindPFlag("tenant", rootCmd.PersistentFlags().Lookup("tenant"))
	viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, _ := os.UserHomeDir()
		viper.AddConfigPath(home)
		viper.SetConfigName(".kuaforia")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("KUAFORIA")
	viper.AutomaticEnv()

	viper.ReadInConfig()
}
