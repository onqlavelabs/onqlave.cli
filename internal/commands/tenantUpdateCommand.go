package commands

import (
	"errors"
	"fmt"
	"github.com/muesli/reflow/wrap"
	"github.com/onqlavelabs/onqlave.core/internal/api"
	"github.com/onqlavelabs/onqlave.core/internal/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
	"os"
	"strings"
)

//func validInput(tenant_name string, tenant_label string) bool {
//	_, err := mail.ParseAddress(address)
//	return err == nil
//}

var tenantLabelUpdate string
var tenantNameUpdate string

func tenantUpdateCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "update",
		Short:   "update.",
		Long:    "update command.",
		Example: "onqlave tenants update",
		Args: func(cmd *cobra.Command, args []string) error {
			//if len(args) < 2 {
			//	return errors.New("requires tenant_name and tenant_label")
			//}
			//
			//if args[0] == "" || args[1] == "" {
			//	return errors.New("requires tenant_name and tenant_label can't be empty")
			//}
			//
			//emailAddress = args[0]
			return nil
		},
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			if !isEnvironmentConfigured() {
				return errors.New("your environment is not configured. please run 'config init' before running any other command")
			}
			if tenantLabelUpdate == "" {
				return fmt.Errorf("Tenant Label can't be empty")
			}

			if tenantNameUpdate == "" {
				return fmt.Errorf("Tenant Name can't be empty")
			}
			return nil
		},
		Run: runTenantUpdateCommand,
	}

	init.Flags().StringVarP(&tenantLabelUpdate, "tenant_label", "l", "", "Update Tenant Label")
	init.Flags().StringVarP(&tenantNameUpdate, "tenant_name", "n", "", "Update Tenant Name")
	return init
}

func runTenantUpdateCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	apiService := api.NewAPIIntegrationService(api.APIIntegrationServiceOptions{Ctx: cmd.Context()})

	_, err := apiService.UpdateTenant(tenantNameUpdate, tenantLabelUpdate)
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error describing tenant '%s': %s", viper.GetString("tenant_id"), err)) + "\n")
		return
	}
	s := &strings.Builder{}
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Green).Padding(1, 1, 1, 0).Render(wrap.String("Tenant Updated Successful", width)))
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String("====================", width)))
	fmt.Println(s.String())
}
