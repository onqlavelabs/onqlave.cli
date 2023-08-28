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

type setDefaultProjectOperation struct {
	projectId               string
	projectOperationTimeout int
}

var setDefaultProject setDefaultProjectOperation

func setDefaultCommand() *cobra.Command {
	setDefaultProject.projectOperationTimeout = 10
	return &cobra.Command{
		Use:     "default",
		Short:   "set default project by ID",
		Long:    "This command is used to set default project by ID. Project id is required.",
		Example: "onqlave project default",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Project id is required")))
			}
			setDefaultProject.projectId = args[0]
			return nil
		},
		Run: runSetDefaultCommand,
	}
}

func runSetDefaultCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	projectID := setDefaultProject.projectId

	_, err := newProjectAPIService(cmd.Context()).SetDefaultProject(projectID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error setting default project '%s': ", projectID), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Setting default project sometime takes up to %d minutes.", setDefaultProject.projectOperationTimeout)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	common.CliRenderSuccessActionResourceOutput(width, projectID, common.ResourceProject, common.ActionSetDefault)
}
