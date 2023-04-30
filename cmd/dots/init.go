package dots

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/matt-riley/dots-cli/internal/config"
	"github.com/matt-riley/dots-cli/internal/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initialize() *cobra.Command {
	init := &cobra.Command{
		Use:     "init",
		Short:   "initialize the config file",
		Long:    "creates a config file in the specified directory",
		Example: "mattd init\n mattd i",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			viper.AutomaticEnv()
			viper.SetEnvPrefix("mattd")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			exists := config.FileExists()
			spin := tea.NewProgram(tui.InitialInitModel(exists), tea.WithAltScreen())
			if err := spin.Start(); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}
	return init
}
