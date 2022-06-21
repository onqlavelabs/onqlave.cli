package commands

import (
	"fmt"

	"github.com/TylerBrock/colorjson"
	"github.com/onqlavelabs/onqlave.core/internal/api"
	"github.com/onqlavelabs/onqlave.core/internal/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func tenantDescribeCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "describe",
		Short:   "describe.",
		Long:    "describe command.",
		Example: "onqlave tenant describe",
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		Run: runTenantDiscoverCommand,
	}
	return init
}

func runTenantDiscoverCommand(cmd *cobra.Command, args []string) {
	//width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	apiService := api.NewAPIIntegrationService(api.APIIntegrationServiceOptions{Ctx: cmd.Context()})

	tenant, err := apiService.GetTenant()
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error describing tenant '%s': %s", viper.GetString("tenant_id"), err)) + "\n")
		return
	}

	// s := &strings.Builder{}
	// header := fmt.Sprintf("Tenant Name: '%s'.", tenant.Name)
	// s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	// s.WriteString("\n")
	// fmt.Println(s.String())
	s, _ := colorjson.Marshal(tenant)
	fmt.Println(string(s))
}
