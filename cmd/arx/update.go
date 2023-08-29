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

type updateProjectOperation struct {
	projectId               string
	projectName             string
	projectRegion           string
	projectOperationTimeout int
	projectSpendLimit       uint64
	projectRotationCycle    string
	projectOwner            string
	projectIsDefault        bool
}

func (o updateProjectOperation) waitForCompletion(apiService *arx.Service, projectId string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	message := "Waiting for provider update to complete."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})
	for duration.Minutes() < float64(valid) {
		result, err := apiService.CheckProjectOperationState(projectId, arx.UpdateOperation)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		}
		time.Sleep(time.Millisecond * 50)

	}
}

var updateProject updateProjectOperation

func updateCommand() *cobra.Command {
	updateProject.projectOperationTimeout = 10
	init := &cobra.Command{
		Use:   "update",
		Short: "update provider by ID and attributes",
		Long: "This command is used to update provider by ID. Project id, provider name, " +
			"provider region, provider encryption method, provider rotation cycle, provider owner, provider spend limit " +
			"and provider is default are required.",
		Example: "onqlave provider update",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return cliCommon.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Project id is required")))
			}
			updateProject.projectId = args[0]
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

			projectApiService := newProjectAPIService(cmd.Context())

			modelWrapper, err := projectApiService.GetProjectBaseInfo()
			if err != nil {
				return cliCommon.CliRenderErr(cmd, err)
			}

			baseInfo := projectApiService.GetProjectBaseInfoIDSlice(modelWrapper)

			projectDetail, err := projectApiService.GetProjectDetail(updateProject.projectId)
			if err != nil {
				return cliCommon.CliRenderErr(cmd, err)
			}

			if projectDetail == nil {
				return cliCommon.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Project detail is required")))
			}

			_, err = projectApiService.ValidateEditProjectRequest(
				baseInfo,
				projectDetail.ProviderID,
				updateProject.projectRegion,
				updateProject.projectRotationCycle,
			)
			if err != nil {
				return cliCommon.CliRenderErr(cmd, err)
			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runUpdateCommand,
	}
	init.Flags().StringVarP(&updateProject.projectName, "project_name", "n", "test", "enter project name")
	init.Flags().StringVarP(&updateProject.projectRegion, "project_region", "r", "", "enter project region - (AUS-EAST, AUS-WEST)")
	init.Flags().StringVarP(&updateProject.projectRotationCycle, "project_rotation_cycle", "c", "Default", "enter project rotation cycle")
	init.Flags().StringVarP(&updateProject.projectOwner, "project_owner", "o", "Default", "enter project owner")
	init.Flags().Uint64VarP(&updateProject.projectSpendLimit, "project_spend_limit", "l", 0, "enter project spend limit")
	init.Flags().BoolVarP(&updateProject.projectIsDefault, "project_is_default", "i", false, "enter project is default")

	return init
}

func runUpdateCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	projectID := updateProject.projectId

	projectApiService := newProjectAPIService(cmd.Context())
	_, err := projectApiService.UpdateProject(contracts.UpdateArx{
		ID:            common.ArxId(projectID),
		Name:          updateProject.projectName,
		Regions:       []string{updateProject.projectRegion},
		RotationCycle: updateProject.projectRotationCycle,
		Owner:         updateProject.projectOwner,
		SpendLimit:    utils.UInt64(updateProject.projectSpendLimit),
		IsDefault:     utils.Bool(updateProject.projectIsDefault),
	})
	if err != nil {
		cliCommon.RenderCLIOutputError(fmt.Sprintf("There was an error updating provider '%s': ", projectID), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Provider update sometime takes up to %d minutes.", updateProject.projectOperationTimeout)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	// Run the function.
	ui, err := utils.NewSpinnerTUI(cmd.Context(), utils.SpinnerOptions{
		Valid:    cliCommon.Valid,
		Consumer: communication.GetConsumer(),
	})
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up provider update operation: %s", err)) + "\n")
		return
	}
	go updateProject.waitForCompletion(projectApiService, projectID, communication.GetProducer(), updateProject.projectOperationTimeout)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up provider update operation: %s", err)) + "\n")
		return
	}

	cliCommon.CliRenderUIErrorOutput(ui, cliCommon.ResourceProject, cliCommon.ActionUpdated, projectID)
}
