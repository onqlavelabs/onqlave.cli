package user

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api/user"
)

func Command() *cobra.Command {
	userCmd := &cobra.Command{
		Use:                "user",
		Short:              "user management",
		Long:               "This command is used to manage users resources.",
		Example:            "onqlave user",
		PersistentPreRunE:  common.PersistentPreRun,
		PersistentPostRunE: common.PersistentPostRunE,
	}

	userCmd.AddCommand(listCommand())

	return userCmd
}

func newUserApiService(ctx context.Context) *user.Service {
	return user.NewService(user.ServiceOpt{Ctx: ctx})
}
