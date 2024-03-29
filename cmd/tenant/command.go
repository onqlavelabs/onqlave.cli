package tenant

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api"
)

func Command() *cobra.Command {
	tenantCmd := &cobra.Command{
		Use:               "tenant",
		Short:             "tenant management",
		Long:              "This command is used to manage tenants resource.",
		Example:           "onqlave tenant",
		PersistentPreRunE: common.PersistentPreRunE,
		PersistentPostRun: common.PersistentPostRun,
	}

	tenantCmd.AddCommand(
		describeCommand(),
		updateCommand(),
	)

	return tenantCmd
}

func newTenantApiService(ctx context.Context) *api.APIIntegrationService {
	return api.NewAPIIntegrationService(api.APIIntegrationServiceOptions{Ctx: ctx})
}
