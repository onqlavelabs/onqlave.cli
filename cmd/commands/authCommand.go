package commands

import (
	"github.com/spf13/cobra"
)

func authCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "auth",
		Short:   "auth.",
		Long:    "auth parent command.",
		Example: "onqlave auth",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	init.AddCommand(signupCommand(), loginCommand())
	return init
}
