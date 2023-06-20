package application

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/core/errors"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
)

type archiveApplicationOperation struct {
	applicationId string
}

var _archiveApplication archiveApplicationOperation

func archiveCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "archive",
		Short:   "archive application by ID",
		Long:    "This command is used to archive application by ID. Application ID is required.",
		Example: "onqlave application archive",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("ApplicationID is required")))
			}
			_archiveApplication.applicationId = args[0]
			return nil
		},
		Run: runArchiveCommand,
	}
}

func runArchiveCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	AppID := _archiveApplication.applicationId

	err := newApplicationAPIService(cmd.Context()).ArchiveApplication(AppID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error archive application '%s': ", AppID), err)
		return
	}

	common.CliRenderSuccessActionResourceOutput(width, AppID, common.ResourceApplication, common.ActionArchived)
}
