// Package dots the main dots cmd package
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
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

// Execute executes the root command.
func Execute() error {
	rootCmd.AddCommand(initialize())
	return rootCmd.ExecuteContext(context.Background())
}
