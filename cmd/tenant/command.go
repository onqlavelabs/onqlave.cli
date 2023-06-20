package tenant

import (
	"context"
	"github.com/onqlavelabs/onqlave.cli/internal/cli/api"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
)

func Command() *cobra.Command {
	tenantCmd := &cobra.Command{
		Use:               "tenant",
		Short:             "tenant management",
		Long:              "This command is used to manage tenants resource.",
		Example:           "onqlave tenant",
		PersistentPreRunE: common.PersistentPreRun,
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
