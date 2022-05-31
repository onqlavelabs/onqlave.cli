package commands

import (
	"errors"
	"fmt"

	"github.com/onqlavelabs/onqlave.core/internal/cli"
	"github.com/onqlavelabs/onqlave.core/internal/configs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configPath string
	env        string
)

func configInitCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "init",
		Short:   "init.",
		Long:    "init command.",
		Example: "onqlave configs init",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			if env != "dev" && env != "prod" {
				return errors.New("'env' flag is invalid. It should be either 'dev' or 'prod'")
			}
			return nil
		},
		RunE: runConfigInitCommand,
	}
	init.PersistentFlags().StringVarP(&configPath, cfgPath, "c", GetConfigFilePath(), "Location of the onqlave config file")
	init.Flags().StringVarP(&env, environment, "e", "dev", "Environment (dev/prod)")
	return init
}

func runConfigInitCommand(cmd *cobra.Command, args []string) error {
	err := configs.CreateFile(configPath)
	if err != nil {
		return err
	}
	fmt.Println(cli.BoldStyle.Copy().Foreground(cli.Green).Render("🎉 Done!  You successfully initialize your environment . Next step is to signup/login is you already haven't.\n"))
	fmt.Println(cli.TextStyle.Render("For more information, read our documentation at https://www.onqlave.com/docs\n"))
	return nil
}
