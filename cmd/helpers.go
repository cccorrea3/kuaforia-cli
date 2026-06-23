package cmd

import (
	"strconv"
	"strings"

	"github.com/kuaforia/cli/internal/client"
	"github.com/kuaforia/cli/internal/format"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getClient() *client.Client {
	return client.New(
		viper.GetString("server"),
		viper.GetString("api_key"),
		viper.GetString("tenant"),
	)
}

func getWorkspace() int {
	return viper.GetInt("default_workspace")
}

func getFormatter() format.Formatter {
	return format.New(format.FormatType(viper.GetString("output")))
}

func parseIntArg(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func splitFlag(cmd *cobra.Command, name string) []string {
	v := cmd.Flag(name).Value.String()
	if v == "" {
		return nil
	}
	parts := strings.Split(v, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
