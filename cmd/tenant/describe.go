package tenant

import (
	"encoding/json"
	"fmt"
	"github.com/onqlavelabs/onqlave.cli/internal/cli/api"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
)

func describeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "describe",
		Short:   "describe tenant",
		Long:    "This command is used to describe tenant info. It will return information of your current user's tenant.",
		Example: "onqlave tenant describe",
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: runDescribeCommand,
	}
}

func runDescribeCommand(cmd *cobra.Command, args []string) {
	var tenant api.TenantInfo
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	tenantDetail, err := newTenantApiService(cmd.Context()).GetTenant()
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error describing tenant '%s': ", viper.GetString("tenant_id")), err)
		return
	}

	tenantBytes, err := json.Marshal(tenantDetail["data"])
	if err != nil {
		return
	}
	err = json.Unmarshal(tenantBytes, &tenant)
	if err != nil {
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderDescribeResourceOutput(width, tenantDetail, common.ResourceTenant, viper.GetString("tenant_id"))
		return
	}

	common.NewDataTable(tenant).Render()
}
