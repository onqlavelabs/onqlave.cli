package commands

import "github.com/spf13/cobra"

func vaultsCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "vaults",
		Short:   "vaults.",
		Long:    "vaults parent command.",
		Example: "onqlave vaults",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	init.AddCommand(addVaultCommand())
	init.AddCommand(vaultListCommand())
	return init
}
