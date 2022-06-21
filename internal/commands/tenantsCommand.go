package commands

import "github.com/spf13/cobra"

func tenantsCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "tenants",
		Short:   "tenants",
		Long:    "tenants",
		Example: "onqlave tenants",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	init.AddCommand(tenantDescribeCommand())
	return init
}
