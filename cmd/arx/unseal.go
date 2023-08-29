package arx

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wrap"
	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api"
	"github.com/onqlavelabs/onqlave.cli/internal/api/arx"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type unsealProjectOperation struct {
	projectId               string
	projectOperationTimeout int
}

func (o unsealProjectOperation) waitForCompletion(apiService *arx.Service, projectId string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	message := "Waiting for project unseal to complete."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})
	for duration.Minutes() < float64(valid) {
		result, err := apiService.CheckProjectOperationState(projectId, arx.UnsealOperation)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		}
		time.Sleep(time.Millisecond * 50)
	}
}

var unsealProject unsealProjectOperation

func unsealCommand() *cobra.Command {
	unsealProject.projectOperationTimeout = 10
	return &cobra.Command{
		Use:     "unseal",
		Short:   "unseal project by ID",
		Long:    "This command is used to unseal project by ID. Project id is required.",
		Example: "onqlave project unseal",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Project id is required")))
			}
			unsealProject.projectId = args[0]
			return nil
		},
		Run: runUnsealCommand,
	}
}

func runUnsealCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	projectID := unsealProject.projectId

	projectApiService := newProjectAPIService(cmd.Context())

	projectDetail, err := projectApiService.GetProjectDetail(projectID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error getting project detail '%s': ", projectID), err)
		return
	}

	if projectDetail.Acl.CanNot["unseal_reason"] != "" {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error unseal project: %s", projectDetail.Acl.CanNot["unseal_reason"])))
		return
	}

	projectId, err := projectApiService.UnsealProject(projectID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error retry unsealing project '%s': ", projectID), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Project unseal sometime takes up to %d minutes.", unsealProject.projectOperationTimeout)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	ui, err := utils.NewSpinnerTUI(cmd.Context(), utils.SpinnerOptions{
		Valid:    common.Valid,
		Consumer: communication.GetConsumer(),
	})
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up project unseal operation: %s", err)) + "\n")
		return
	}
	go unsealProject.waitForCompletion(projectApiService, projectId, communication.GetProducer(), unsealProject.projectOperationTimeout)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up project unseal operation: %s", err)) + "\n")
		return
	}

	common.CliRenderUIErrorOutput(ui, common.ResourceProject, common.ActionUnsealed, projectID)
}
