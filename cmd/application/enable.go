package application

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/core/errors"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/cli"
)

type enableApplicationOperation struct {
	applicationId string
}

var _enableApplication enableApplicationOperation

func enableCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "enable",
		Short:   "enable application by ID",
		Long:    "This command is used to enable application by ID. Application ID is required.",
		Example: "onqlave application enable",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, cli.BoldStyle.Render("ApplicationID is required")))
			}
			_enableApplication.applicationId = args[0]
			return nil
		},
		Run: runEnableCommand,
	}
}

func runEnableCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	AppID := _enableApplication.applicationId

	err := newApplicationAPIService(cmd.Context()).EnableApplication(AppID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error enable application '%s': ", AppID), err)
		return
	}

	common.CliRenderSuccessActionResourceOutput(width, AppID, common.ResourceApplication, common.ActionEnabled)

}
