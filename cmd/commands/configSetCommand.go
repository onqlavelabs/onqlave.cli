package commands

import (
	"errors"
	"fmt"

	"github.com/onqlavelabs/onqlave.core/internal/cli"
	"github.com/onqlavelabs/onqlave.core/internal/configs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func configSetCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "set",
		Short:   "set.",
		Long:    "set command.",
		Example: "onqlave configs set",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			if env != "dev" && env != "prod" {
				return errors.New("'env' flag is invalid. It should be either 'dev' or 'prod'")
			}
			return nil
		},
		RunE: runConfigSetCommand,
	}
	init.Flags().StringVarP(&env, environment, "e", "dev", "Environment (dev/prod)")
	return init
}

func runConfigSetCommand(cmd *cobra.Command, args []string) error {
	err := configs.CreateFile(configPath)
	if err != nil {
		return err
	}
	fmt.Println(cli.BoldStyle.Copy().Foreground(cli.Green).Render("🎉 Done!  You successfully set your environment .\n"))
	fmt.Println(cli.TextStyle.Render("For more information, read our documentation at https://www.onqlave.com/docs\n"))
	return nil
}
