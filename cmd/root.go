package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/cmd/application"
	"github.com/onqlavelabs/onqlave.cli/cmd/arx"
	"github.com/onqlavelabs/onqlave.cli/cmd/auth"
	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/cmd/config"
	"github.com/onqlavelabs/onqlave.cli/cmd/key"
	"github.com/onqlavelabs/onqlave.cli/cmd/tenant"
	"github.com/onqlavelabs/onqlave.cli/cmd/user"
)

var rootCmd = &cobra.Command{
	Version:           common.Version,
	Use:               "onqlave CLI helps you manage your Onqlave environment.",
	Example:           "onqlave",
	PersistentPreRunE: RootPreRunE,
}

func Execute() {
	initConfig()
	addCommands()

	err := rootCmd.ExecuteContext(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	rootCmd.PersistentFlags().Bool(common.FlagJson, false, "JSON Output. Set to true if stdout is not a TTY.")
	viper.SetDefault(common.FlagApiBaseUrl, "")
	viper.AddConfigPath(common.GetConfigDir())
	viper.SetConfigName(common.ConfigFile)     // Register config file name (no extension)
	viper.SetConfigType(common.ConfigTypeJson) // Look for specific type

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func addCommands() {
	rootCmd.AddCommand(config.Command())
	rootCmd.AddCommand(auth.Command())
	rootCmd.AddCommand(key.Command())
	rootCmd.AddCommand(tenant.Command())
	rootCmd.AddCommand(arx.Command())
	rootCmd.AddCommand(application.Command())
	rootCmd.AddCommand(user.Command())
}

func RootPreRunE(cmd *cobra.Command, args []string) error {
	if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
		cmd.SilenceUsage = true
		return common.ReplacePersistentPreRunE(cmd, err)
	}
	return nil
}
