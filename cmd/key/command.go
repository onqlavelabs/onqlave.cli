package key

import (
	"context"
	"github.com/onqlavelabs/onqlave.cli/internal/cli/api/apiKey"

	"github.com/spf13/cobra"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
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
		Use:               "key",
		Short:             "api key management",
		Long:              "This command is used to manage api key resources.",
		Example:           "onqlave key",
		PersistentPreRunE: common.PersistentPreRun,
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

func newKeyApiService(ctx context.Context) *apiKey.APIKeyIntegrationService {
	return apiKey.NewAPIKeyIntegrationService(apiKey.APIKeyIntegrationServiceOptions{Ctx: ctx})
}
