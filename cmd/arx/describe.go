package arx

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

type describeProjectOperation struct {
	projectId string
}

var describeProject describeProjectOperation

func describeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "describe",
		Short:   "describe project by ID",
		Long:    "This command is used to describe project by ID. Project id is required.",
		Example: "onqlave project describe",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Project id is required")))
			}
			describeProject.projectId = args[0]
			return nil
		},
		Run: runDescribeCommand,
	}
}

func runDescribeCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	projectID := describeProject.projectId

	projectDetail, err := newProjectAPIService(cmd.Context()).GetProjectDetail(projectID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error describing project '%s': ", projectID), err)
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderDescribeResourceOutput(width, projectDetail.Detail, common.ResourceProject, projectID)
		return
	}

	common.NewDataTable(projectDetail.Detail).Render()
}
