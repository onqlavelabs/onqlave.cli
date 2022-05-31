package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func appDisableCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "disable",
		Short:   "disable.",
		Long:    "disable command.",
		Example: "onqlave apps disable",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: runAppDisableCommand,
	}
	return init
}

func runAppDisableCommand(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented yet")
}
