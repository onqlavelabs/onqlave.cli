package arx

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type describeArxOperation struct {
	arxId string
}

var _describeArx describeArxOperation

func describeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "describe",
		Short:   "describe arx by ID",
		Long:    "This command is used to describe arx by ID. Arx id is required.",
		Example: "onqlave arx describe",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("ArxID is required")))
			}
			_describeArx.arxId = args[0]
			return nil
		},
		Run: runDescribeCommand,
	}
}

func runDescribeCommand(cmd *cobra.Command, args []string) {
	if viper.GetBool(common.FlagDebug) {
		start := time.Now()
		defer func() {
			fmt.Printf("Took: %s\n", time.Since(start))
		}()
	}
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	ArxID := _describeArx.arxId

	arxDetail, err := newArxAPIService(cmd.Context()).GetArxDetail(ArxID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error describing arx '%s': ", ArxID), err)
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderDescribeResourceOutput(width, arxDetail.Detail, common.ResourceArx, ArxID)
		return
	}

	common.NewDataTable(arxDetail.Detail).Render()
}
