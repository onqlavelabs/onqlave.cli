package tenant

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/errors"
)

var tenantLabelUpdate string
var tenantNameUpdate string

func updateCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "update",
		Short:   "update tenant by name and label",
		Long:    "This command is used to update tenant. Tenant name and tenant label are required.",
		Example: "onqlave tenants update",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if tenantLabelUpdate == "" && tenantNameUpdate == "" {
				return common.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Tenant label and tenant name can not be both empty")))
			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runUpdateCommand,
	}

	init.Flags().StringVarP(&tenantLabelUpdate, "tenant_label", "l", "", "update tenant label")
	init.Flags().StringVarP(&tenantNameUpdate, "tenant_name", "n", "", "update tenant name")
	return init
}

func runUpdateCommand(cmd *cobra.Command, args []string) {
	var tenantMap map[string]interface{}

	apiService := newTenantApiService(cmd.Context())
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	if tenantNameUpdate == "" || tenantLabelUpdate == "" {
		tenant, err := apiService.GetTenant()
		if err != nil {
			common.RenderCLIOutputError(fmt.Sprintf("There was an error describing tenant '%s': ", viper.GetString("tenant_id")), err)
			return
		}
		tenantMap = tenant["data"].(map[string]interface{})
	}

	if tenantLabelUpdate == "" {
		tenantLabelUpdate = tenantMap["tenant_label"].(string)
	}
	if tenantNameUpdate == "" {
		tenantNameUpdate = tenantMap["tenant_name"].(string)
	}

	_, err := apiService.UpdateTenant(tenantNameUpdate, tenantLabelUpdate)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error describing tenant '%s': ", viper.GetString("tenant_id")), err)
		return
	}

	common.CliRenderSuccessActionResourceOutput(width, viper.GetString("tenant_id"), common.ResourceTenant, "update")
}
