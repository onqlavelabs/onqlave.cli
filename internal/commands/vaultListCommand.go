package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	onlyEnableVaults bool
)

func vaultListCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "list",
		Short:   "list.",
		Long:    "list command.",
		Example: "onqlave vaults list",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: runVaultListCommand,
	}
	init.Flags().BoolVarP(&onlyEnableVaults, "only-enabled", "o", true, "Only list the enable vaults")
	return init
}

func runVaultListCommand(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented yet")
}
