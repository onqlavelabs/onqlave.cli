package application

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

type describeApplicationOperation struct {
	applicationId string
}

var describeApplication describeApplicationOperation

func describeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "describe",
		Short:   "describe application by ID",
		Long:    "This command is used to describe application by ID. Application ID is required.",
		Example: "onqlave application describe",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("ApplicationID is required")))
			}
			describeApplication.applicationId = args[0]
			return nil
		},
		Run: runDescribeCommand,
	}
}

func runDescribeCommand(cmd *cobra.Command, args []string) {
	AppID := describeApplication.applicationId
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	applicationDetail, err := newApplicationAPIService(cmd.Context()).GetApplication(AppID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error describing application '%s': ", AppID), err)
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderDescribeResourceOutput(width, applicationDetail, common.ResourceApplication, AppID)
		return
	}

	common.NewDataTable(applicationDetail).Render()
}
