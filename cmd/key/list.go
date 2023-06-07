package key

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
		Short:   "list api keys",
		Long:    "This command is used to list all existing api keys.",
		Example: "onqlave key list",
		Run:     runListCommand,
	}
}

func runListCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	data, err := newKeyApiService(cmd.Context()).GetKeys()
	if err != nil {
		common.RenderCLIOutputError("There was an error getting api keys: ", err)
		return
	}

	if len(data.Keys) == 0 {
		common.CliRenderListResourceOutputNoRecord(width)
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderListResourceOutput(width, data.Keys, common.ResourceKey)
		return
	}

	common.NewDataTable(data.Keys).Render()
}
