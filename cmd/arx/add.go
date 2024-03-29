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
	contracts "github.com/onqlavelabs/onqlave.core/contracts/arx"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type addArxOperation struct {
	arxName             string
	arxType             string
	arxProvider         string
	arxPurpose          string
	arxRegion           string
	arxDescription      string
	arxOperationTimeout int
	arxSpendLimit       uint64
	arxEncryptionMethod string
	arxRotationCycle    string
	arxOwner            string
	arxIsDefault        bool
}

func (o addArxOperation) waitForCompletion(apiService *arx.Service, arxId string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	message := "Waiting for arx creation to complete."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})
	for duration.Minutes() < float64(valid) {
		result, err := apiService.CheckArxOperationState(arxId, arx.AddOperation)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		}
		time.Sleep(time.Millisecond * 50)
	}
}

var addArx addArxOperation

func addCommand() *cobra.Command {
	addArx.arxOperationTimeout = 10
	init := &cobra.Command{
		Use:   "add",
		Short: "add arx by name and attributes",
		Long: "This command is used to add arx. Valid arx name, arx provider,  arx type, arx purpose, " +
			"arx region, arx description, arx encryption method, arx rotation cycle, arx owner, arx spend limit " +
			"and arx is default are required.",
		Example: "onqlave arx add",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Arx name is required")))
			}
			addArx.arxName = args[0]
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return common.CliRenderErr(cmd, err)
			}

			if !common.IsLoggedIn() {
				return common.CliRenderErr(cmd, common.ErrRequireLogIn)
			}

			arxApiService := newArxAPIService(cmd.Context())
			modelWrapper, err := arxApiService.GetArxBaseInfo()
			if err != nil {
				return common.CliRenderErr(cmd, err)
			}

			baseInfo := arxApiService.GetArxBaseInfoIDSlice(modelWrapper)
			_, err = arxApiService.ValidateArx(
				baseInfo,
				addArx.arxProvider,
				addArx.arxType,
				addArx.arxPurpose,
				addArx.arxRegion,
				addArx.arxEncryptionMethod,
				addArx.arxRotationCycle,
				addArx.arxOwner,
			)
			if err != nil {
				return common.CliRenderErr(cmd, err)
			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runAddCommand,
	}
	init.Flags().StringVarP(&addArx.arxProvider, "arx_provider", "p", "", "enter arx cloud provider")
	init.Flags().StringVarP(&addArx.arxType, "arx_type", "t", "", "enter arx type")
	init.Flags().StringVarP(&addArx.arxPurpose, "arx_purpose", "u", "", "enter arx purpose")
	init.Flags().StringVarP(&addArx.arxRegion, "arx_region", "r", "", "enter arx region")
	init.Flags().StringVarP(&addArx.arxDescription, "arx_description", "d", "", "enter arx description")
	init.Flags().StringVarP(&addArx.arxEncryptionMethod, "arx_encryption_method", "e", "", "enter arx encryption method")
	init.Flags().StringVarP(&addArx.arxRotationCycle, "arx_rotation_cycle", "c", "", "enter arx rotation cycle")
	init.Flags().StringVarP(&addArx.arxOwner, "arx_owner", "o", "", "enter arx owner")
	init.Flags().Uint64VarP(&addArx.arxSpendLimit, "arx_spend_limit", "l", 0, "enter arx spend limit")
	init.Flags().BoolVarP(&addArx.arxIsDefault, "arx_is_default", "i", false, "enter arx is default")

	return init
}

func runAddCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	arxApiService := newArxAPIService(cmd.Context())
	regions := strings.Split(addArx.arxRegion, ",")
	arxId, err := arxApiService.AddArx(contracts.NewArx{
		Name:             addArx.arxName,
		Plan:             addArx.arxType,
		Provider:         addArx.arxProvider,
		Purpose:          addArx.arxPurpose,
		Regions:          regions,
		EncryptionMethod: addArx.arxEncryptionMethod,
		RotationCycle:    addArx.arxRotationCycle,
		Owner:            addArx.arxOwner,
		SpendLimit:       utils.UInt64(addArx.arxSpendLimit),
		IsDefault:        addArx.arxIsDefault,
	})
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error creating arx '%s': ", addArx.arxName), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Arx creation sometime takes up to %d minutes.", addArx.arxOperationTimeout)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	ui, err := utils.NewSpinnerTUI(cmd.Context(), utils.SpinnerOptions{
		Valid:    common.Valid,
		Consumer: communication.GetConsumer(),
	})
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up arx creation operation: %s", err)) + "\n")
		return
	}
	go addArx.waitForCompletion(arxApiService, arxId, communication.GetProducer(), addArx.arxOperationTimeout)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up arx creation operation: %s", err)) + "\n")
		return
	}

	common.CliRenderUIErrorOutput(ui, common.ResourceArx, common.ActionCreated, arxId)
}
