package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func appDescribeCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "describe",
		Short:   "describe.",
		Long:    "describe command.",
		Example: "onqlave apps describe",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: runAppDescribeCommand,
	}
	return init
}

func runAppDescribeCommand(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented yet")
}
