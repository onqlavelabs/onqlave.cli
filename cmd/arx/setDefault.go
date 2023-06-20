package arx

import (
	"fmt"
	"os"
	"strings"

	"github.com/muesli/reflow/wrap"
	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/core/errors"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/cli"
)

type setDefaultArxOperation struct {
	arxId               string
	arxOperationTimeout int
}

var _setDefaultArx setDefaultArxOperation

func setDefaultCommand() *cobra.Command {
	_setDefaultArx.arxOperationTimeout = 10
	return &cobra.Command{
		Use:     "default",
		Short:   "set default arx by ID",
		Long:    "This command is used to set default arx by ID. Arx id is required.",
		Example: "onqlave arx default",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				cmd.SilenceUsage = true

				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, cli.BoldStyle.Render("Arx id is required")))
			}
			_setDefaultArx.arxId = args[0]
			return nil
		},
		Run: runSetDefaultCommand,
	}
}

func runSetDefaultCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	arxID := _setDefaultArx.arxId

	_, err := newArxAPIService(cmd.Context()).SetDefaultArx(arxID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error setting default arx '%s': ", arxID), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Setting default arx sometime takes up to %d minutes.", _setDefaultArx.arxOperationTimeout)
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	common.CliRenderSuccessActionResourceOutput(width, arxID, common.ResourceArx, common.ActionSetDefault)
}
