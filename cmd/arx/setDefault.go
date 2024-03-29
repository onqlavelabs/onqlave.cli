package arx

import (
	"fmt"
	"os"
	"strings"

	"github.com/muesli/reflow/wrap"
	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type setDefaultArxOperation struct {
	arxId               string
	arxOperationTimeout int
}

var setDefaultArx setDefaultArxOperation

func setDefaultCommand() *cobra.Command {
	setDefaultArx.arxOperationTimeout = 10
	return &cobra.Command{
		Use:     "default",
		Short:   "set default arx by ID",
		Long:    "This command is used to set default arx by ID. Arx id is required.",
		Example: "onqlave arx default",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Arx id is required")))
			}
			setDefaultArx.arxId = args[0]
			return nil
		},
		Run: runSetDefaultCommand,
	}
}

func runSetDefaultCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	arxID := setDefaultArx.arxId

	_, err := newArxAPIService(cmd.Context()).SetDefaultArx(arxID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error setting default arx '%s': ", arxID), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Setting default arx sometime takes up to %d minutes.", setDefaultArx.arxOperationTimeout)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	common.CliRenderSuccessActionResourceOutput(width, arxID, common.ResourceArx, common.ActionSetDefault)
}
