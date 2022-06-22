package commands

import (
	"github.com/spf13/cobra"
)

func clustersCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "clusters",
		Short:   "clusters",
		Long:    "clusters",
		Example: "onqlave clusters",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	init.AddCommand(configInitCommand())
	init.AddCommand(configSetCommand())
	init.AddCommand(configListCommand())
	return init
}
