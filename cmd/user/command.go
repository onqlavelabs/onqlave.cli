package user

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api/user"
)

func Command() *cobra.Command {
	userCmd := &cobra.Command{
		Use:               "user",
		Short:             "user management",
		Long:              "This command is used to manage users resources.",
		Example:           "onqlave user",
		PersistentPreRunE: common.PersistentPreRun,
	}

	userCmd.AddCommand(listCommand())

	return userCmd
}

func newUserApiService(ctx context.Context) *user.UserAPIIntegrationService {
	return user.NewUserAPIIntegrationService(user.UserAPIIntegrationServiceOptions{Ctx: ctx})
}
