package commands

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//func validInput(tenant_name string, tenant_label string) bool {
//	_, err := mail.ParseAddress(address)
//	return err == nil
//}

func tenantUpdateCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "update",
		Short:   "update.",
		Long:    "update command.",
		Example: "onqlave tenants update",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("requires tenant_name and tenant_label")
			}

			if args[0] == "" || args[1] == "" {
				return errors.New("requires tenant_name and tenant_label can't be empty")
			}

			emailAddress = args[0]
			return nil
		},
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		Run: runTenantDescribeCommand,
	}
	return init
}
