package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func appUpdateCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "update",
		Short:   "update.",
		Long:    "update command.",
		Example: "onqlave apps update",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: runAppUpdateCommand,
	}
	return init
}

func runAppUpdateCommand(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented yet")
}
