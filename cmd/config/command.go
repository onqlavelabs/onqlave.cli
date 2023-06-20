package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
)

func Command() *cobra.Command {
	configCmd := &cobra.Command{
		Use:     "config",
		Short:   "config environment variables",
		Long:    "This command is used to config your CLI environment variables.",
		Example: "onqlave config",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}

			cmd.SilenceUsage = false

			return nil
		},
	}

	configCmd.AddCommand(
		initCommand(),
		currentCommand(),
	)

	return configCmd
}
