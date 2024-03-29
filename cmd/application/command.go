package application

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api/application"
)

type FlagApplication string

const (
	Tech FlagApplication = "application_technology"
)

func (flag FlagApplication) String() string {
	return string(flag)
}

func Command() *cobra.Command {
	applicationCmd := &cobra.Command{
		Use:               "application",
		Short:             "application management",
		Long:              "This command is used to manage application resources.",
		Example:           "onqlave application",
		PersistentPreRunE: common.PersistentPreRunE,
		PersistentPostRun: common.PersistentPostRun,
	}

	applicationCmd.AddCommand(
		addCommand(),
		updateCommand(),
		listCommand(),
		describeCommand(),
		disableCommand(),
		enableCommand(),
		archiveCommand(),
		baseCommand(),
	)

	return applicationCmd
}

func newApplicationAPIService(ctx context.Context) *application.Service {
	return application.NewService(application.ServiceOpt{Ctx: ctx})
}
