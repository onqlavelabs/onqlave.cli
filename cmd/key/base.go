package key

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/api_key"
)

type BaseKey struct {
	Model      string `json:"model"`
	Flag       string `json:"flag"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Technology string `json:"technology"`
}

func baseCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "base",
		Short:   "get base",
		Long:    "This command is used to get api key base info.",
		Example: "onqlave key base",
		Run:     runBaseCommand,
	}
}

func runBaseCommand(cmd *cobra.Command, args []string) {
	apiService := newKeyApiService(cmd.Context())
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	data, err := apiService.GetKeyBaseInfo()
	if err != nil {
		common.RenderCLIOutputError("There was an error getting api key base information: ", err)
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderBaseResourceOutput(width, data, common.ResourceKey)
		return
	}

	common.NewDataTable(convertKeyBaseInfo(data)).Render()
}

func convertKeyBaseInfo(key api_key.APIKeyModels) []BaseKey {
	var list []BaseKey

	list = append(list, BaseKey{Model: "Applications", Flag: ApplicationID.String()})
	for _, application := range key.Applications {
		list = append(list, BaseKey{
			Id:         application.ID,
			Name:       application.Name,
			Technology: application.ApplicationTechnology.Id,
		})
	}

	list = append(list, BaseKey{Model: "Arx", Flag: ArxID.String()})
	for _, arx := range key.Arx {
		list = append(list, BaseKey{
			Id:         arx.ID,
			Name:       arx.Name,
			Technology: "",
		})
	}

	return list
}
