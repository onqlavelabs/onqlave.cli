package arx

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
)

func listCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "list arx",
		Long:    "This command is used to list all existing arx.",
		Example: "onqlave arx list",
		Run:     runListCommand,
	}
}

func runListCommand(cmd *cobra.Command, args []string) {
	start := time.Now()
	defer common.LogResponseTime(start)

	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	data, err := newArxAPIService(cmd.Context()).GetArx()
	if err != nil {
		common.RenderCLIOutputError("There was an error retry get arx: ", err)
		return
	}

	if len(data.Clusters) == 0 {
		common.CliRenderListResourceOutputNoRecord(width)
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderListResourceOutput(width, data.Clusters, common.ResourceArx)
		return
	}

	common.NewDataTable(data.Clusters).Render()
}
