package auth

import (
	"context"
	"github.com/onqlavelabs/onqlave.cli/internal/cli/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
)

func Command() *cobra.Command {
	authCommand := &cobra.Command{
		Use:     "auth",
		Short:   "authentication",
		Long:    "This command is used to authenticate.",
		Example: "onqlave auth",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}
			return nil
		},
	}

	authCommand.AddCommand(
		loginCommand(),
		signupCommand(),
	)

	return authCommand
}

func newAuthAPIService(ctx context.Context) *api.APIIntegrationService {
	return api.NewAPIIntegrationService(api.APIIntegrationServiceOptions{Ctx: ctx})
}
