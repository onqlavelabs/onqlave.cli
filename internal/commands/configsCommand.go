package commands

import (
	"github.com/spf13/cobra"
)

func configsCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "configs",
		Short:   "configs",
		Long:    "configs",
		Example: "onqlave configs",

		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	init.AddCommand(configInitCommand())
	init.AddCommand(configSetCommand())
	init.AddCommand(configListCommand())
	return init
}
