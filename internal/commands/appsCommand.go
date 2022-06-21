package commands

import (
	"github.com/spf13/cobra"
)

func appsCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "apps",
		Short:   "apps.",
		Long:    "apps parent command.",
		Example: "onqlave apps",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	init.AddCommand(appListCommand())
	init.AddCommand(appAddCommand())
	init.AddCommand(appDescribeCommand())
	init.AddCommand(appDisableCommand())
	init.AddCommand(appEnableCommand())
	init.AddCommand(appUpdateCommand())
	init.AddCommand(appAttachCommand())
	return init
}
