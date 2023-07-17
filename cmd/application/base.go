package application

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.core/contracts/application"
)

type BaseApplication struct {
	Model       string `json:"model"`
	Flag        string `json:"flag"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cors        *bool  `json:"cors"`
	Order       string `json:"order"`
	Icon        string `json:"icon"`
	Enable      *bool  `json:"enable"`
	Default     *bool  `json:"default"`
}

func baseCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "base",
		Short:   "get base application info",
		Long:    "This command is used to get application base info.",
		Example: "onqlave application base",
		Run:     runBaseCommand,
	}
}

func runBaseCommand(cmd *cobra.Command, args []string) {
	if viper.GetBool(common.FlagDebug) {
		fmt.Println(common.DebugStart)
		defer fmt.Println(common.DebugEnd)
	}
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	application, err := newApplicationAPIService(cmd.Context()).GetBaseApplication()
	if err != nil {
		common.RenderCLIOutputError("There was an error getting application base information: ", err)
		return
	}
	if viper.GetBool(common.FlagJson) {
		common.CliRenderBaseResourceOutput(width, application, common.ResourceApplication)
		return
	}

	common.NewDataTable(convertApplicationBaseInfo(application)).Render()
}

func convertApplicationBaseInfo(application application.Technologies) []BaseApplication {
	var list []BaseApplication

	list = append(list, BaseApplication{Model: "Technologies", Flag: Tech.String()})
	for _, tech := range application.Technologies {
		list = append(list, BaseApplication{
			Model:       "",
			ID:          tech.Id,
			Name:        tech.Name,
			Description: tech.Description,
			Cors:        &tech.Cors,
			Order:       fmt.Sprintf("%v", tech.Order),
			Icon:        tech.Icon,
			Enable:      &tech.Enable,
			Default:     &tech.IsDefault,
		})
	}

	return list
}
