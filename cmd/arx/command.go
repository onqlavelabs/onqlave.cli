package arx

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api/arx"
)

type FlagArx string

const (
	Type             FlagArx = "arx_type"
	EncryptionMethod FlagArx = "arx_encryption_method"
	Provider         FlagArx = "arx_provider"
	Purpose          FlagArx = "arx_purpose"
	RotationCycle    FlagArx = "arx_rotation_cycle"
)

func (flag FlagArx) String() string {
	return string(flag)
}

func Command() *cobra.Command {
	arxCmd := &cobra.Command{
		Use:               "arx",
		Short:             "arx management",
		Long:              "This command is used to manage arx resources.",
		Example:           "onqlave arx",
		PersistentPreRunE: common.PersistentPreRunE,
		PersistentPostRun: common.PersistentPostRun,
	}

	arxCmd.AddCommand(
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

	return arxCmd
}

func newArxAPIService(ctx context.Context) *arx.Service {
	return arx.NewService(arx.ServiceOpt{Ctx: ctx})
}
