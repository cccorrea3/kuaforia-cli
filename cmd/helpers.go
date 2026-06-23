package cmd

import (
	"github.com/kuaforia/cli/internal/client"
	"github.com/kuaforia/cli/internal/format"
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
