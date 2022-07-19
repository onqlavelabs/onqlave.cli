package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/muesli/reflow/wrap"
	"github.com/onqlavelabs/onqlave.core/internal/api"
	"github.com/onqlavelabs/onqlave.core/internal/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

func tenantDescribeCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "describe",
		Short:   "describe.",
		Long:    "describe command.",
		Example: "onqlave tenants describe",
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
		Run: runTenantDescribeCommand,
	}
	return init
}

func runTenantDescribeCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	apiService := api.NewAPIIntegrationService(api.APIIntegrationServiceOptions{Ctx: cmd.Context()})

	tenant, err := apiService.GetTenant()
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error describing tenant '%s': %s", viper.GetString("tenant_id"), err)) + "\n")
		return
	}

	s := &strings.Builder{}
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String("Describing your Tenant =>", width)))
	s.WriteString("\n")
	s.WriteString(cli.RenderAsJson(tenant))
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String("====================", width)))
	fmt.Println(s.String())
}
