package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func keysCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "keys",
		Short:   "keys.",
		Long:    "keys parent command.",
		Example: "onqlave keys",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			// Environment variables are expected to be ALL CAPS
			viper.AutomaticEnv()
			viper.SetEnvPrefix(prefix)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// homeDir, err := os.UserHomeDir()
			// if err != nil {
			// 	return err
			// }

			// if err := tea.NewProgram(tui.NewInitPrompt(viper.GetString(cfgPath), homeDir)).Start(); err != nil {
			// 	return err
			// }
			return nil
		},
	}
	// init.AddCommand(appListCommand())
	// init.AddCommand(appAddCommand())
	// init.AddCommand(appDeleteCommand())
	// init.AddCommand(appDescribeCommand())
	// init.AddCommand(appDisableCommand())
	// init.AddCommand(appEnableCommand())
	return init
}
