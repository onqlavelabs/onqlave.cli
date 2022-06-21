package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	newVault Vault
)

type Vault struct {
	name        string
	description string
	encryption  string
	database    string
	custodian   string
}

func addVaultCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "add",
		Short:   "add.",
		Long:    "add command.",
		Example: "onqlave vaults add",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: runVaultAddCommand,
	}
	init.Flags().StringVarP(&newVault.name, "name", "n", "", "Enter a name for your vault")
	init.Flags().StringVarP(&newVault.description, "description", "s", " ", "Enter description for your vault")
	init.Flags().StringVarP(&newVault.encryption, "encryption", "e", "aes-gcm-128", "Specify encryption algorithm used to encrypt data in vault")
	init.Flags().StringVarP(&newVault.database, "database", "d", "postgresql", "Specify backend database used to store data in vault")
	init.Flags().StringVarP(&newVault.custodian, "custodian", "c", "", "Specify the custodian's email address. they would be able to configure the vault")
	return init
}

func runVaultAddCommand(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented yet")
}
