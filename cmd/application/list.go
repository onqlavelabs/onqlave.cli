package application

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
)

func listCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "list application",
		Long:    "This command is used to list all existing applications.",
		Example: "onqlave application list",
		Run:     runListCommand,
	}
}

func runListCommand(cmd *cobra.Command, args []string) {
	if viper.GetBool(common.FlagDebug) {
		fmt.Println(common.DebugStart)
		defer fmt.Println(common.DebugEnd)
	}
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	applications, err := newApplicationAPIService(cmd.Context()).GetApplications()
	if err != nil {
		common.RenderCLIOutputError("There was an error listing application: ", err)
		return
	}

	if len(applications) == 0 {
		common.CliRenderListResourceOutputNoRecord(width)
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderListResourceOutput(width, applications, common.ResourceApplication)
		return
	}

	common.NewDataTable(applications).Render()
}
