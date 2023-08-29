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

type addProjectOperation struct {
	projectName             string
	projectType             string
	projectProvider         string
	projectPurpose          string
	projectRegion           string
	projectDescription      string
	projectOperationTimeout int
	projectSpendLimit       uint64
	projectEncryptionMethod string
	projectRotationCycle    string
	projectOwner            string
	projectIsDefault        bool
}

func (o addProjectOperation) waitForCompletion(apiService *arx.Service, projectId string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	message := "Waiting for project creation to complete."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})
	for duration.Minutes() < float64(valid) {
		result, err := apiService.CheckProjectOperationState(projectId, arx.AddOperation)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		}
		time.Sleep(time.Millisecond * 50)
	}
}

var addProject addProjectOperation

func addCommand() *cobra.Command {
	addProject.projectOperationTimeout = 10
	init := &cobra.Command{
		Use:   "add",
		Short: "add project by name and attributes",
		Long: "This command is used to add project. Valid project name, project project,  project type, project purpose, " +
			"project region, project description, project encryption method, project rotation cycle, project owner, project spend limit " +
			"and project is default are required.",
		Example: "onqlave project add",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Project name is required")))
			}
			addProject.projectName = args[0]
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return common.CliRenderErr(cmd, err)
			}

			if !common.IsLoggedIn() {
				return common.CliRenderErr(cmd, common.ErrRequireLogIn)
			}

			projectApiService := newProjectAPIService(cmd.Context())
			modelWrapper, err := projectApiService.GetProjectBaseInfo()
			if err != nil {
				return common.CliRenderErr(cmd, err)
			}

			baseInfo := projectApiService.GetProjectBaseInfoIDSlice(modelWrapper)
			_, err = projectApiService.ValidateProject(
				baseInfo,
				addProject.projectProvider,
				addProject.projectType,
				addProject.projectPurpose,
				addProject.projectRegion,
				addProject.projectEncryptionMethod,
				addProject.projectRotationCycle,
				addProject.projectOwner,
			)
			if err != nil {
				return common.CliRenderErr(cmd, err)
			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runAddCommand,
	}
	init.Flags().StringVarP(&addProject.projectProvider, "project_provider", "p", "", "enter project cloud project")
	init.Flags().StringVarP(&addProject.projectType, "project_type", "t", "", "enter project type")
	init.Flags().StringVarP(&addProject.projectPurpose, "project_purpose", "u", "", "enter project purpose")
	init.Flags().StringVarP(&addProject.projectRegion, "project_region", "r", "", "enter project region")
	init.Flags().StringVarP(&addProject.projectDescription, "project_description", "d", "", "enter project description")
	init.Flags().StringVarP(&addProject.projectEncryptionMethod, "project_encryption_method", "e", "", "enter project encryption method")
	init.Flags().StringVarP(&addProject.projectRotationCycle, "project_rotation_cycle", "c", "", "enter project rotation cycle")
	init.Flags().StringVarP(&addProject.projectOwner, "project_owner", "o", "", "enter project owner")
	init.Flags().Uint64VarP(&addProject.projectSpendLimit, "project_spend_limit", "l", 0, "enter project spend limit")
	init.Flags().BoolVarP(&addProject.projectIsDefault, "project_is_default", "i", false, "enter project is default")

	return init
}

func runAddCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	projectApiService := newProjectAPIService(cmd.Context())
	regions := strings.Split(addProject.projectRegion, ",")
	projectId, err := projectApiService.AddProject(contracts.NewArx{
		Name:             addProject.projectName,
		Plan:             addProject.projectType,
		Provider:         addProject.projectProvider,
		Purpose:          addProject.projectPurpose,
		Regions:          regions,
		EncryptionMethod: addProject.projectEncryptionMethod,
		RotationCycle:    addProject.projectRotationCycle,
		Owner:            addProject.projectOwner,
		SpendLimit:       utils.UInt64(addProject.projectSpendLimit),
		IsDefault:        addProject.projectIsDefault,
	})
	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error creating project '%s': ", addProject.projectName), err)
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Project creation sometime takes up to %d minutes.", addProject.projectOperationTimeout)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	ui, err := utils.NewSpinnerTUI(cmd.Context(), utils.SpinnerOptions{
		Valid:    common.Valid,
		Consumer: communication.GetConsumer(),
	})
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up project creation operation: %s", err)) + "\n")
		return
	}
	go addProject.waitForCompletion(projectApiService, projectId, communication.GetProducer(), addProject.projectOperationTimeout)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up project creation operation: %s", err)) + "\n")
		return
	}

	common.CliRenderUIErrorOutput(ui, common.ResourceProject, common.ActionCreated, projectId)
}
