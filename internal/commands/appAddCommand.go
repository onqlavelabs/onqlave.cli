package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func appAddCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "add",
		Short:   "add.",
		Long:    "add command.",
		Example: "onqlave apps add",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: runAppAddCommand,
	}
	init.Flags().StringVarP(&emailAddress, "email-address", "e", "", "Enter your email address to signup")
	return init
}

func runAppAddCommand(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented yet")
}
