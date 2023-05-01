package dots

import (
	"context"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "mattd",
	Short:   "mattd is a dotfile manager",
	Long:    "mattd is a dotfile manager. It helps you manage your dotfiles by storing them in a git repository and symlinking them to your home directory.",
	Example: "mattd",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() error {
	return rootCmd.ExecuteContext(context.Background())
}
