package application

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type disableApplicationOperation struct {
	applicationId string
}

var _disableApplication disableApplicationOperation

func disableCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "disable",
		Short:   "disable application by ID",
		Long:    "This command is used to disable application by ID. Application ID is required.",
		Example: "onqlave application disable",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("ApplicationID is required")))
			}
			_disableApplication.applicationId = args[0]
			return nil
		},
		Run: runDisableCommand,
	}
}

func runDisableCommand(cmd *cobra.Command, args []string) {
	if viper.GetBool(common.FlagDebug) {
		fmt.Println(common.DebugStart)
		defer fmt.Println(common.DebugEnd)
	}
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	AppID := _disableApplication.applicationId

	err := newApplicationAPIService(cmd.Context()).DisableApplication(AppID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error disable application '%s': ", AppID), err)
		return
	}

	common.CliRenderSuccessActionResourceOutput(width, AppID, common.ResourceApplication, common.ActionDisabled)
}
