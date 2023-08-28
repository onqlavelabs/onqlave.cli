package arx

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
)

func listCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "list project",
		Long:    "This command is used to list all existing project.",
		Example: "onqlave project list",
		Run:     runListCommand,
	}
}

func runListCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	data, err := newProjectAPIService(cmd.Context()).GetProject()
	if err != nil {
		common.RenderCLIOutputError("There was an error retry get project: ", err)
		return
	}

	if len(data.Projects) == 0 {
		common.CliRenderListResourceOutputNoRecord(width)
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderListResourceOutput(width, data.Projects, common.ResourceProject)
		return
	}

	common.NewDataTable(data.Projects).Render()
}
