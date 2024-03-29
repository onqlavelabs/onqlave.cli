package key

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type describeKeyOperation struct {
	keyId string
}

var describeKey describeKeyOperation

func describeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "describe",
		Short:   "describe api key by ID",
		Long:    "This command is used to describe api key detail by ID. Api key ID is required.",
		Example: "onqlave key describe",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("KeyID is required")))
			}
			describeKey.keyId = args[0]
			return nil
		},
		Run: runDescribeCommand,
	}
}

func runDescribeCommand(cmd *cobra.Command, args []string) {
	keyId := describeKey.keyId
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	data, err := newKeyApiService(cmd.Context()).GetKeyDetail(keyId)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error describing api key '%s': ", keyId), err)
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderDescribeResourceOutput(width, data, common.ResourceKey, keyId)
		return
	}

	common.NewDataTable(data).Render()
}
