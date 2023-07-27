package auth

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api"
)

func Command() *cobra.Command {
	authCommand := &cobra.Command{
		Use:               "auth",
		Short:             "authentication",
		Long:              "This command is used to authenticate.",
		Example:           "onqlave auth",
		PersistentPreRunE: common.PersistentPreRunE,
		PersistentPostRun: common.PersistentPostRun,
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
