package arx

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wrap"
	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api/arx"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/cli"
)

type deleteArxOperation struct {
	arxId               string
	arxOperationTimeout int
}

func (o deleteArxOperation) waitForCompletion(apiService *arx.ArxAPIIntegrationService, arxId string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	message := "Waiting for arx deletion to complete."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})
	for duration.Minutes() < float64(valid) {
		result, err := apiService.CheckArxOperationState(arxId, arx.DeleteOperation)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		}
		time.Sleep(time.Millisecond * 50)
	}
}

var _deleteArx deleteArxOperation

func deleteCommand() *cobra.Command {
	_deleteArx.arxOperationTimeout = 10
	return &cobra.Command{
		Use:     "delete",
		Short:   "delete arx by ID",
		Long:    "This command is used to delete arx by ID. Arx id is required.",
		Example: "onqlave arx delete",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				cmd.SilenceUsage = true

				return errors.New("arx id is required")
			}
			_deleteArx.arxId = args[0]
			return nil
		},
		Run: runDeleteCommand,
	}
}

func runDeleteCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	arxID := _deleteArx.arxId

	arxApiService := newArxAPIService(cmd.Context())
	_, err := arxApiService.DeleteArx(arxID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error retry deleting arx '%s': ", arxID), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Arx deletion sometime takes up to %d minutes.", _deleteArx.arxOperationTimeout)
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	// Run the function.
	ui, err := cli.NewSpnnerTUI(cmd.Context(), cli.SpinnerOptions{
		Valid:    common.Valid,
		Consumer: communication.GetConsumer(),
	})
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up arx delete operation: %s", err)) + "\n")
		return
	}
	go _deleteArx.waitForCompletion(arxApiService, arxID, communication.GetProducer(), _deleteArx.arxOperationTimeout)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up arx delete operation: %s", err)) + "\n")
		return
	}

	common.CliRenderUIErrorOutput(ui, common.ResourceArx, common.ActionDeleted, arxID)
}
