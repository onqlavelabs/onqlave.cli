package key

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wrap"
	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api"
	"github.com/onqlavelabs/onqlave.cli/internal/api/apiKey"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type deleteAPIKeyOperation struct {
	keyID            string
	operationTimeout int
}

var _deleteAPIKey deleteAPIKeyOperation

func deleteCommand() *cobra.Command {
	_deleteAPIKey.operationTimeout = 10
	return &cobra.Command{
		Use:     "delete",
		Short:   "delete api key by ID",
		Long:    "This command is used to delete api key by ID. Api key ID is required.",
		Example: "onqlave key delete",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("KeyID is required")))
			}

			_deleteAPIKey.keyID = args[0]
			return nil
		},
		Run: runDeleteCommand,
	}
}

func runDeleteCommand(cmd *cobra.Command, args []string) {
	if viper.GetBool(common.FlagDebug) {
		fmt.Println(common.DebugStart)
		defer fmt.Println(common.DebugEnd)
	}
	deleteKeyID := _deleteAPIKey.keyID
	apiService := newKeyApiService(cmd.Context())
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	keyID, err := apiService.DeleteKey(deleteKeyID)
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error deleting api key '%s': ", deleteKeyID), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Api key deletion sometime takes up to %d minutes.", _deleteAPIKey.operationTimeout)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	ui, err := utils.NewSpnnerTUI(cmd.Context(), utils.SpinnerOptions{Valid: common.Valid, Consumer: communication.GetConsumer()})
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up api key delete operation: %s", err)) + "\n")
		return
	}

	go _deleteAPIKey.waitForCompletion(apiService, keyID, communication.GetProducer(), _deleteAPIKey.operationTimeout)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up api key delete operation: %s", err)) + "\n")
		return
	}

	common.CliRenderUIErrorOutput(ui, common.ResourceKey, common.ActionDeleted, keyID)
}

func (o deleteAPIKeyOperation) waitForCompletion(apiService *apiKey.Service, keyId string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	message := "Waiting for api key deletion completion."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})

	for duration.Minutes() < float64(valid) {
		result, err := apiService.CheckAPIKeyOperationStatus(keyId, apiKey.DeleteOperation)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		}
		time.Sleep(time.Millisecond * 50)
	}
}
