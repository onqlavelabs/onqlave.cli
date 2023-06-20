package arx

import (
	"context"
	"github.com/onqlavelabs/onqlave.cli/internal/cli/api/arx"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
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
		PersistentPreRunE: common.PersistentPreRun,
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

func newArxAPIService(ctx context.Context) *arx.ArxAPIIntegrationService {
	return arx.NewArxAPIIntegrationService(arx.ArxAPIIntegrationServiceOptions{
		Ctx: ctx,
	})
}
