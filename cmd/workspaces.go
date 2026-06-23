package cmd

import (
	"github.com/spf13/cobra"
)

var workspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "Gestionar workspaces",
}

var workspacesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar workspaces del tenant",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("list_workspaces", map[string]any{})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

var workspacesGetCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Detalle de workspace",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("get_workspace", map[string]any{
			"workspaceId": parseIntArg(args[0]),
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

var workspacesMembersCmd = &cobra.Command{
	Use:   "members [id]",
	Short: "Miembros del workspace",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("get_workspace_members", map[string]any{
			"workspaceId": parseIntArg(args[0]),
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	workspacesCmd.AddCommand(workspacesListCmd)
	workspacesCmd.AddCommand(workspacesGetCmd)
	workspacesCmd.AddCommand(workspacesMembersCmd)
	rootCmd.AddCommand(workspacesCmd)
}
