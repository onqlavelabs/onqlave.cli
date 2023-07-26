package arx

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	contracts "github.com/onqlavelabs/onqlave.core/contracts/arx"
)

type BaseArx struct {
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
		Short:   "get base arx info",
		Long:    "This command is used to get arx base info.",
		Example: "onqlave arx base",
		Run:     runBaseCommand,
	}
}

func runBaseCommand(cmd *cobra.Command, args []string) {
	start := time.Now()
	defer common.LogResponseTime(start)

	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	arx, err := newArxAPIService(cmd.Context()).GetArxBaseInfo()
	if err != nil {
		common.RenderCLIOutputError("There was an error getting arx base information: ", err)
		return
	}
	if viper.GetBool(common.FlagJson) {
		common.CliRenderBaseResourceOutput(width, arx, common.ResourceArx)
		return
	}

	common.NewDataTable(convertArxBaseInfo(arx)).Render()
}

func convertArxBaseInfo(arx contracts.BaseInfo) []BaseArx {
	var list []BaseArx

	list = append(list, BaseArx{Model: "Plan", Flag: Type.String()})
	for _, plan := range arx.Plans {
		list = append(list, BaseArx{
			ID:      plan.ID,
			Name:    plan.Name,
			Default: plan.IsDefault,
			Enable:  plan.Enable,
			Order:   fmt.Sprintf("%v", *plan.Order),
		})
	}

	list = append(list, BaseArx{Model: "Encryption Method", Flag: EncryptionMethod.String()})
	for _, method := range arx.EncryptionMethods {
		list = append(list, BaseArx{
			ID:      method.ID,
			Name:    method.Name,
			Default: method.IsDefault,
			Enable:  method.Enable,
			Order:   fmt.Sprintf("%v", *method.Order),
		})
	}

	list = append(list, BaseArx{Model: "Provider", Flag: Provider.String()})
	for _, provider := range arx.Providers {
		var regions []string

		for _, region := range provider.Regions {
			regions = append(regions, region.ID)
		}

		list = append(list, BaseArx{
			ID:      provider.ID,
			Name:    provider.Name,
			Default: provider.IsDefault,
			Enable:  provider.Enable,
			Order:   fmt.Sprintf("%v", *provider.Order),
			Regions: strings.Join(regions, ","),
		})
	}

	list = append(list, BaseArx{Model: "Purpose", Flag: Purpose.String()})
	for _, purpose := range arx.Purposes {
		list = append(list, BaseArx{
			ID:      purpose.ID,
			Name:    purpose.Name,
			Default: purpose.IsDefault,
			Enable:  purpose.Enable,
			Order:   fmt.Sprintf("%v", *purpose.Order),
		})
	}

	list = append(list, BaseArx{Model: "Rotation Cycle", Flag: RotationCycle.String()})
	for _, cycle := range arx.RotationCycles {
		list = append(list, BaseArx{
			ID:      cycle.ID,
			Name:    cycle.Name,
			Default: cycle.IsDefault,
			Enable:  cycle.Enable,
			Order:   fmt.Sprintf("%v", *cycle.Order),
		})
	}

	return list
}
