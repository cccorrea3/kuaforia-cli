package cmd

import (
	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Gestionar usuarios del workspace",
}

var usersListCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar usuarios",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("list_users", map[string]any{})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

var usersGetCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Detalle de usuario",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("get_user", map[string]any{
			"userId": parseIntArg(args[0]),
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

var usersInviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invitar un usuario al workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		email, _ := cmd.Flags().GetString("email")
		role, _ := cmd.Flags().GetString("role")
		workspace, _ := cmd.Flags().GetInt("workspace")

		result, err := getClient().Call("invite_user", map[string]any{
			"email":       email,
			"role":        role,
			"workspaceId": workspace,
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

var usersDisableCmd = &cobra.Command{
	Use:   "disable [id]",
	Short: "Deshabilitar usuario",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getClient().Call("disable_user", map[string]any{
			"userId": parseIntArg(args[0]),
		})
		if err != nil {
			return err
		}
		getFormatter().JSON(result)
		return nil
	},
}

func init() {
	usersInviteCmd.Flags().String("email", "", "Email del usuario")
	usersInviteCmd.Flags().String("role", "analyst", "Rol: analyst, editor, admin")
	usersInviteCmd.Flags().Int("workspace", 0, "ID del workspace")

	usersCmd.AddCommand(usersListCmd)
	usersCmd.AddCommand(usersGetCmd)
	usersCmd.AddCommand(usersInviteCmd)
	usersCmd.AddCommand(usersDisableCmd)
	rootCmd.AddCommand(usersCmd)
}
