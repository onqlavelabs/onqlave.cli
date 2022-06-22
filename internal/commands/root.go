package commands

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Execute is the command line applications entry function
func Execute() error {
	rootCmd := &cobra.Command{
		Version: version,
		Use:     "onqlave (onqlave) is a CLI that helps you manage your Onqlave environment.",
		Example: "onqlave",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
				return err
			}
			// viper.AutomaticEnv()
			// viper.SetEnvPrefix(prefix)

			// dir, err := os.UserHomeDir()
			// if err != nil {
			// 	return err
			// }

			// viper.AddConfigPath(dir + cfgDir)
			// viper.SetConfigName(cfgFile) // Register config file name (no extension)
			// viper.SetConfigType("yaml")  // Look for specific type
			// viper.ReadInConfig()
			return nil
		},
	}

	// rootCmd.SetHelpFunc(boa.HelpFunc)
	// rootCmd.SetUsageFunc(boa.UsageFunc)
	rootCmd.PersistentFlags().Bool("json", false, "Output logs as JSON.  Set to true if stdout is not a TTY.")

	// Add sub commands
	rootCmd.AddCommand(configsCommand())
	rootCmd.AddCommand(authCommand())
	rootCmd.AddCommand(vaultsCommand())
	rootCmd.AddCommand(appsCommand())
	rootCmd.AddCommand(keysCommand())
	rootCmd.AddCommand(tenantsCommand())
	rootCmd.AddCommand(clustersCommand())

	viper.SetDefault("api_base_url", "")
	viper.AddConfigPath(GetConfigDir())
	viper.SetConfigName(cfgFile) // Register config file name (no extension)
	viper.SetConfigType("json")  // Look for specific type
	viper.ReadInConfig()

	return rootCmd.ExecuteContext(context.Background())
}
