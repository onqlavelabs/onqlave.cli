package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	onlyEnabled bool
)

func appListCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "list",
		Short:   "list.",
		Long:    "list command.",
		Example: "onqlave apps list",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: runAppListCommand,
	}
	init.Flags().BoolVarP(&onlyEnabled, "only-enabled", "o", true, "Enter your email address to signup")
	return init
}

func runAppListCommand(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented yet")
}
