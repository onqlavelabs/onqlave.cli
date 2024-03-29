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

	cliCommon "github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api"
	"github.com/onqlavelabs/onqlave.cli/internal/api/arx"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	contracts "github.com/onqlavelabs/onqlave.core/contracts/arx"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type updateArxOperation struct {
	arxId               string
	arxName             string
	arxRegion           string
	arxOperationTimeout int
	arxSpendLimit       uint64
	arxRotationCycle    string
	arxOwner            string
	arxIsDefault        bool
}

func (o updateArxOperation) waitForCompletion(apiService *arx.Service, arxId string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	message := "Waiting for arx update to complete."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})
	for duration.Minutes() < float64(valid) {
		result, err := apiService.CheckArxOperationState(arxId, arx.UpdateOperation)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		}
		time.Sleep(time.Millisecond * 50)

	}
}

var updateArx updateArxOperation

func updateCommand() *cobra.Command {
	updateArx.arxOperationTimeout = 10
	init := &cobra.Command{
		Use:   "update",
		Short: "update arx by ID and attributes",
		Long: "This command is used to update arx by ID. Arx id, arx name, " +
			"arx region, arx encryption method, arx rotation cycle, arx owner, arx spend limit " +
			"and arx is default are required.",
		Example: "onqlave arx update",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return cliCommon.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("ArxID is required")))
			}
			updateArx.arxId = args[0]
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return cliCommon.CliRenderErr(cmd, err)
			}

			if !cliCommon.IsLoggedIn() {
				return cliCommon.CliRenderErr(cmd, cliCommon.ErrRequireLogIn)
			}

			arxApiService := newArxAPIService(cmd.Context())

			modelWrapper, err := arxApiService.GetArxBaseInfo()
			if err != nil {
				return cliCommon.CliRenderErr(cmd, err)
			}

			baseInfo := arxApiService.GetArxBaseInfoIDSlice(modelWrapper)

			arxDetail, err := arxApiService.GetArxDetail(updateArx.arxId)
			if err != nil {
				return cliCommon.CliRenderErr(cmd, err)
			}

			if arxDetail == nil {
				return cliCommon.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Arx detail is required")))
			}

			_, err = arxApiService.ValidateEditArxRequest(
				baseInfo,
				arxDetail.ProviderID,
				updateArx.arxRegion,
				updateArx.arxRotationCycle,
			)
			if err != nil {
				return cliCommon.CliRenderErr(cmd, err)
			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runArxUpdateCommand,
	}
	init.Flags().StringVarP(&updateArx.arxName, "arx_name", "n", "test", "enter arx name")
	init.Flags().StringVarP(&updateArx.arxRegion, "arx_region", "r", "", "enter arx region - (AUS-EAST, AUS-WEST)")
	init.Flags().StringVarP(&updateArx.arxRotationCycle, "arx_rotation_cycle", "c", "Default", "enter arx rotation cycle")
	init.Flags().StringVarP(&updateArx.arxOwner, "arx_owner", "o", "Default", "enter arx owner")
	init.Flags().Uint64VarP(&updateArx.arxSpendLimit, "arx_spend_limit", "l", 0, "enter arx spend limit")
	init.Flags().BoolVarP(&updateArx.arxIsDefault, "arx_is_default", "i", false, "enter arx is default")

	return init
}

func runArxUpdateCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	arxID := updateArx.arxId

	arxApiService := newArxAPIService(cmd.Context())
	arxId, err := arxApiService.UpdateArx(contracts.UpdateArx{
		ID:            common.ArxId(arxID),
		Name:          updateArx.arxName,
		Regions:       []string{updateArx.arxRegion},
		RotationCycle: updateArx.arxRotationCycle,
		Owner:         updateArx.arxOwner,
		SpendLimit:    utils.UInt64(updateArx.arxSpendLimit),
		IsDefault:     utils.Bool(updateArx.arxIsDefault),
	})
	if err != nil {
		cliCommon.RenderCLIOutputError(fmt.Sprintf("There was an error updating arx '%s': ", arxID), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Arx update sometime takes up to %d minutes.", updateArx.arxOperationTimeout)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	// Run the function.
	ui, err := utils.NewSpinnerTUI(cmd.Context(), utils.SpinnerOptions{
		Valid:    cliCommon.Valid,
		Consumer: communication.GetConsumer(),
	})
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up arx update operation: %s", err)) + "\n")
		return
	}
	go updateArx.waitForCompletion(arxApiService, arxId, communication.GetProducer(), updateArx.arxOperationTimeout)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up arx update operation: %s", err)) + "\n")
		return
	}

	cliCommon.CliRenderUIErrorOutput(ui, cliCommon.ResourceArx, cliCommon.ActionUpdated, arxID)
}
