package key

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api/apiKey"
)

type FlagApiKey string

const (
	ArxID         FlagApiKey = "key_arx_id"
	ApplicationID FlagApiKey = "key_application_id"
)

func (flag FlagApiKey) String() string {
	return string(flag)
}

func Command() *cobra.Command {
	keyCmd := &cobra.Command{
		Use:                "key",
		Short:              "api key management",
		Long:               "This command is used to manage api key resources.",
		Example:            "onqlave key",
		PersistentPreRunE:  common.PersistentPreRun,
		PersistentPostRunE: common.PersistentPostRunE,
	}

	keyCmd.AddCommand(
		addCommand(),
		deleteCommand(),
		describeCommand(),
		listCommand(),
		baseCommand(),
	)

	return keyCmd
}

func newKeyApiService(ctx context.Context) *apiKey.Service {
	return apiKey.NewService(apiKey.ServiceOpt{Ctx: ctx})
}
