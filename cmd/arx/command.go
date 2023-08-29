package arx

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api/arx"
)

type FlagProject string

const (
	Type             FlagProject = "project_type"
	EncryptionMethod FlagProject = "project_encryption_method"
	Provider         FlagProject = "project_provider"
	Purpose          FlagProject = "project_purpose"
	RotationCycle    FlagProject = "project_rotation_cycle"
)

func (flag FlagProject) String() string {
	return string(flag)
}

func Command() *cobra.Command {
	projectCmd := &cobra.Command{
		Use:               "project",
		Short:             "project management",
		Long:              "This command is used to manage project resources.",
		Example:           "onqlave project",
		PersistentPreRunE: common.PersistentPreRunE,
		PersistentPostRun: common.PersistentPostRun,
	}

	projectCmd.AddCommand(
		addCommand(),
		deleteCommand(),
		describeCommand(),
		updateCommand(),
		listCommand(),
		sealCommand(),
		setDefaultCommand(),
		unsealCommand(),
		baseCommand(),
	)

	return projectCmd
}

func newProjectAPIService(ctx context.Context) *arx.Service {
	return arx.NewService(arx.ServiceOpt{Ctx: ctx})
}
