package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func configListCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "list",
		Short:   "list.",
		Long:    "list command.",
		Example: "onqlave configs list",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: runConfigListCommand,
	}
	return init
}

func runConfigListCommand(cmd *cobra.Command, args []string) error {
	return errors.New("it is not implemented yet")
	// fmt.Println(cli.TextStyle.Render("For more information, read our documentation at https://www.onqlave.com/docs\n"))
	// return nil
}
