package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func loginCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "login",
		Short:   "login.",
		Long:    "login command.",
		Example: "onqlave auth login",
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			if !validMailAddress(emailAddress) {
				return errors.New("email address is invalid. Please provide a valid email address")
			}
			return nil
		},
		RunE: _runConfigLoginCommand,
	}
	init.Flags().StringVarP(&emailAddress, "email-address", "e", "", "Enter your email address to signup")
	return init
}

func _runConfigLoginCommand(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented yet")
	// tenantService := api.NewTenantService(api.TenantServiceOptions{Ctx: cmd.Context()})

	// link, err := tenantService.GetLink()
	// if err != nil {
	// 	return err
	// }
	// // Run the function.
	// ui, err := cli.NewSignupTUI(cmd.Context(), cli.SignupOptions{
	// 	Link:  link,
	// 	Ctx:   cmd.Context(),
	// 	Valid: 10,
	// })
	// if err != nil {
	// 	return err
	// }
	// if err := tea.NewProgram(ui).Start(); err != nil {
	// 	return err
	// }
	// // So we can exit with a non-zero code.
	// return ui.Error()
	//return nil
}
