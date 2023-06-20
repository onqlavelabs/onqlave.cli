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
	coreErr "github.com/onqlavelabs/onqlave.cli/core/errors"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api/arx"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/cli"
)

type unsealArxOperation struct {
	arxId               string
	arxOperationTimeout int
}

func (o unsealArxOperation) waitForCompletion(apiService *arx.ArxAPIIntegrationService, arxId string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	message := "Waiting for arx unseal to complete."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})
	for duration.Minutes() < float64(valid) {
		result, err := apiService.CheckArxOperationState(arxId, arx.UnsealOperation)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		}
		time.Sleep(time.Millisecond * 50)
	}
}

var _unsealArx unsealArxOperation

func unsealCommand() *cobra.Command {
	_unsealArx.arxOperationTimeout = 10
	return &cobra.Command{
		Use:     "unseal",
		Short:   "unseal arx by ID",
		Long:    "This command is used to unseal arx by ID. Arx id is required.",
		Example: "onqlave arx unseal",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				cmd.SilenceUsage = true

				return common.ReplacePersistentPreRunE(cmd, coreErr.NewCLIResultError(coreErr.KeyCLIMissingRequiredField, cli.BoldStyle.Render("ArxID is required")))
			}
			_unsealArx.arxId = args[0]
			return nil
		},
		Run: runUnsealCommand,
	}
}

func runUnsealCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	arxID := _unsealArx.arxId

	arxApiService := newArxAPIService(cmd.Context())

	arxDetail, err := arxApiService.GetArxDetail(arxID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error getting arx detail '%s': ", arxID), err)
		return
	}

	if arxDetail.Acl.CanNot["unseal_reason"] != "" {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error unseal arx: %s", arxDetail.Acl.CanNot["unseal_reason"])))
		return
	}

	arxId, err := arxApiService.UnsealArx(arxID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error retry unsealing arx '%s': ", arxID), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Arx unseal sometime takes up to %d minutes.", _unsealArx.arxOperationTimeout)
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	// Run the function.
	ui, err := cli.NewSpnnerTUI(cmd.Context(), cli.SpinnerOptions{
		Valid:    common.Valid,
		Consumer: communication.GetConsumer(),
	})
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up arx unseal operation: %s", err)) + "\n")
		return
	}
	go _unsealArx.waitForCompletion(arxApiService, arxId, communication.GetProducer(), _unsealArx.arxOperationTimeout)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up arx unseal operation: %s", err)) + "\n")
		return
	}

	common.CliRenderUIErrorOutput(ui, common.ResourceArx, common.ActionUnsealed, arxID)
}
