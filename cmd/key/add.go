package key

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
	"github.com/onqlavelabs/onqlave.cli/core/contracts/api_key"
	"github.com/onqlavelabs/onqlave.cli/internal/api"
	"github.com/onqlavelabs/onqlave.cli/internal/api/apiKey"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
)

type addApiKeyOperation struct {
	applicationID         string
	arxID                 string
	applicationTechnology string
	operationTimeout      int
}

var _addApiKeyOperation addApiKeyOperation

func addCommand() *cobra.Command {
	_addApiKeyOperation.operationTimeout = 10
	init := &cobra.Command{
		Use:     "add",
		Short:   "add api key by attributes",
		Long:    "This command is used to create api key. Key application ID, arx ID and application technology is required.",
		Example: "onqlave key add",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}
			if !common.IsLoggedIn() {
				return common.ReplacePersistentPreRunE(cmd, common.ErrRequireLogIn)
			}

			apiService := newKeyApiService(cmd.Context())
			baseInfo, err := apiService.GetKeyBaseInfo()
			if err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}

			_, err = apiService.ValidateAPIKey(baseInfo,
				_addApiKeyOperation.applicationID,
				_addApiKeyOperation.arxID,
				_addApiKeyOperation.applicationTechnology,
			)
			if err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runAddCommand,
	}

	init.Flags().StringVarP(&_addApiKeyOperation.applicationID, "key_application_id", "a", "", "enter application id")
	init.Flags().StringVarP(&_addApiKeyOperation.arxID, "key_arx_id", "c", "", "enter arx id")
	init.Flags().StringVarP(&_addApiKeyOperation.applicationTechnology, "key_application_technology", "t", "", "enter application technology")

	return init
}

func runAddCommand(cmd *cobra.Command, args []string) {
	apiService := newKeyApiService(cmd.Context())
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	keyID, err := apiService.AddKey(api_key.CreateAPIKey{
		ApplicationID:         _addApiKeyOperation.applicationID,
		ClusterID:             _addApiKeyOperation.arxID,
		ApplicationTechnology: _addApiKeyOperation.applicationTechnology,
	})
	if err != nil {
		common.RenderCLIOutputError("There was an error creating api key: ", err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Api key creation sometime takes up to %d minutes.", _addApiKeyOperation.operationTimeout)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	ui, err := utils.NewSpnnerTUI(cmd.Context(), utils.SpinnerOptions{Valid: common.Valid, Consumer: communication.GetConsumer()})
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up api key creation operation: %s", err)) + "\n")
		return
	}

	go _addApiKeyOperation.waitForCompletion(apiService, keyID, communication.GetProducer(), _addApiKeyOperation.operationTimeout)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up api key creation operation: %s", err)) + "\n")
		return
	}

	common.CliRenderUIErrorOutput(ui, common.ResourceKey, common.ActionCreated, keyID)
}

func (o addApiKeyOperation) waitForCompletion(apiService *apiKey.Service, keyId string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	message := "Waiting for api key creation completion."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})

	for duration.Minutes() < float64(valid) {
		result, err := apiService.CheckAPIKeyOperationStatus(keyId, apiKey.AddOperation)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		}
		time.Sleep(time.Millisecond * 50)
	}
}
