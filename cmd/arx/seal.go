package arx

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wrap"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api"
	"github.com/onqlavelabs/onqlave.cli/internal/api/arx"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type sealArxOperation struct {
	arxId               string
	arxOperationTimeout int
}

func (o sealArxOperation) waitForCompletion(apiService *arx.Service, arxId string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	message := "Waiting for arx seal to complete."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})
	for duration.Minutes() < float64(valid) {
		result, err := apiService.CheckArxOperationState(arxId, arx.SealOperation)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		}
		time.Sleep(time.Millisecond * 50)
	}
}

var _sealArx sealArxOperation

func sealCommand() *cobra.Command {
	_sealArx.arxOperationTimeout = 10
	return &cobra.Command{
		Use:     "seal",
		Short:   "seal arx by ID",
		Long:    "This command is used to seal arx by ID. Arx id is required.",
		Example: "onqlave arx seal",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("ArxID is required")))
			}
			_sealArx.arxId = args[0]
			return nil
		},
		Run: runSealCommand,
	}
}

func runSealCommand(cmd *cobra.Command, args []string) {
	if viper.GetBool(common.FlagDebug) {
		start := time.Now()
		defer func() {
			fmt.Printf("Took: %s\n", time.Since(start))
		}()
	}
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	arxID := _sealArx.arxId

	arxApiService := newArxAPIService(cmd.Context())
	arxId, err := arxApiService.SealArx(arxID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error retry sealing arx '%s': ", arxID), err)
		return
	}
	s := &strings.Builder{}
	header := fmt.Sprintf("Arx seal sometime takes up to %d minutes.", _sealArx.arxOperationTimeout)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	ui, err := utils.NewSpnnerTUI(cmd.Context(), utils.SpinnerOptions{
		Valid:    common.Valid,
		Consumer: communication.GetConsumer(),
	})
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up arx seal operation: %s", err)) + "\n")
		return
	}
	go _sealArx.waitForCompletion(arxApiService, arxId, communication.GetProducer(), _sealArx.arxOperationTimeout)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up arx seal operation: %s", err)) + "\n")
		return
	}
	common.CliRenderUIErrorOutput(ui, common.ResourceArx, common.ActionSealed, arxID)

}
