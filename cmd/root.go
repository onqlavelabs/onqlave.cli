package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/cmd/arx"
	"github.com/onqlavelabs/onqlave.cli/cmd/auth"
	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/cmd/config"
	"github.com/onqlavelabs/onqlave.cli/cmd/key"
	"github.com/onqlavelabs/onqlave.cli/cmd/tenant"
	"github.com/onqlavelabs/onqlave.cli/cmd/user"
)

var Version = "alpha"
var rootCmd = &cobra.Command{
	Version:           Version,
	Use:               "onqlave CLI helps you manage your Onqlave environment.",
	Example:           "onqlave",
	PersistentPreRunE: RootPreRunE,
}

func Execute() {
	initConfig()
	addCommands()

	err := rootCmd.ExecuteContext(context.Background())
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	rootCmd.SetVersionTemplate("{{.Version}}\n")
	rootCmd.PersistentFlags().Bool(common.FlagJson, false, "JSON Output. Set to true if stdout is not a TTY.")
	rootCmd.PersistentFlags().Bool(common.FlagDebug, false, "Debug mode. Set true for debugging.")

	viper.SetDefault(common.FlagApiBaseUrl, "")
	viper.AddConfigPath(common.GetConfigDir())
	viper.SetConfigName(common.ConfigFile)     // Register config file name (no extension)
	viper.SetConfigType(common.ConfigTypeJson) // Look for specific type

	_ = viper.ReadInConfig()
}

func addCommands() {
	rootCmd.AddCommand(config.Command())
	rootCmd.AddCommand(auth.Command())
	rootCmd.AddCommand(key.Command())
	rootCmd.AddCommand(tenant.Command())
	rootCmd.AddCommand(arx.Command())
	rootCmd.AddCommand(user.Command())
}

func RootPreRunE(cmd *cobra.Command, args []string) error {
	if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
		return common.CliRenderErr(cmd, err)
	}

	return nil
}
