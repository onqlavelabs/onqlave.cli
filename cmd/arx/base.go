package arx

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	contracts "github.com/onqlavelabs/onqlave.core/contracts/arx"
)

type BaseProject struct {
	Model   string `json:"model"`
	Flag    string `json:"flag"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Default *bool  `json:"default"`
	Enable  *bool  `json:"enable"`
	Order   string `json:"order"`
	Regions string `json:"regions"`
}

func baseCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "base",
		Short:   "get base project info",
		Long:    "This command is used to get project base info.",
		Example: "onqlave project base",
		Run:     runBaseCommand,
	}
}

func runBaseCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	project, err := newProjectAPIService(cmd.Context()).GetProjectBaseInfo()
	if err != nil {
		common.RenderCLIOutputError("There was an error getting project base information: ", err)
		return
	}
	if viper.GetBool(common.FlagJson) {
		common.CliRenderBaseResourceOutput(width, project, common.ResourceProject)
		return
	}

	common.NewDataTable(convertProjectBaseInfo(project)).Render()
}

func convertProjectBaseInfo(project contracts.BaseInfo) []BaseProject {
	var list []BaseProject

	list = append(list, BaseProject{Model: "Plan", Flag: Type.String()})
	for _, plan := range project.Plans {
		list = append(list, BaseProject{
			ID:      plan.ID,
			Name:    plan.Name,
			Default: plan.IsDefault,
			Enable:  plan.Enable,
			Order:   fmt.Sprintf("%v", *plan.Order),
		})
	}

	list = append(list, BaseProject{Model: "Encryption Method", Flag: EncryptionMethod.String()})
	for _, method := range project.EncryptionMethods {
		list = append(list, BaseProject{
			ID:      method.ID,
			Name:    method.Name,
			Default: method.IsDefault,
			Enable:  method.Enable,
			Order:   fmt.Sprintf("%v", *method.Order),
		})
	}

	list = append(list, BaseProject{Model: "Provider", Flag: Provider.String()})
	for _, provider := range project.Providers {
		var regions []string

		for _, region := range provider.Regions {
			regions = append(regions, region.ID)
		}

		list = append(list, BaseProject{
			ID:      provider.ID,
			Name:    provider.Name,
			Default: provider.IsDefault,
			Enable:  provider.Enable,
			Order:   fmt.Sprintf("%v", *provider.Order),
			Regions: strings.Join(regions, ","),
		})
	}

	list = append(list, BaseProject{Model: "Purpose", Flag: Purpose.String()})
	for _, purpose := range project.Purposes {
		list = append(list, BaseProject{
			ID:      purpose.ID,
			Name:    purpose.Name,
			Default: purpose.IsDefault,
			Enable:  purpose.Enable,
			Order:   fmt.Sprintf("%v", *purpose.Order),
		})
	}

	list = append(list, BaseProject{Model: "Rotation Cycle", Flag: RotationCycle.String()})
	for _, cycle := range project.RotationCycles {
		list = append(list, BaseProject{
			ID:      cycle.ID,
			Name:    cycle.Name,
			Default: cycle.IsDefault,
			Enable:  cycle.Enable,
			Order:   fmt.Sprintf("%v", *cycle.Order),
		})
	}

	return list
}
