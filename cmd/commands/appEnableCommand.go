package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func appEnableCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "enable",
		Short:   "enable.",
		Long:    "enable command.",
		Example: "onqlave apps enable",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: runAppEnableCommand,
	}
	return init
}

func runAppEnableCommand(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented yet")
}
