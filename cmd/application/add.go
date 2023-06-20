package application

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	contractApplication "github.com/onqlavelabs/onqlave.cli/core/contracts/application"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api/application"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api/user"
)

type addApplicationOperation struct {
	applicationName        string
	applicationDescription string
	applicationTechnology  string
	applicationOwner       string
	applicationCors        string
}

var _addApplicationOperation addApplicationOperation

func addCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "add",
		Short:   "add application by name and attributes",
		Long:    "This command is used to add application by name. Application name, application technology and application owner are required.",
		Example: "onqlave application add",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				cmd.SilenceUsage = true
				return errors.New("application name is required")
			}
			_addApplicationOperation.applicationName = args[0]
			return nil
		},
		// used to overwrite/skip the parent commands persistentPreRunE func
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			apiService := application.NewApplicationAPIIntegrationService(application.ApplicationAPIIntegrationServiceOptions{Ctx: cmd.Context()})

			modelWrapper, err := apiService.GetBaseApplication()
			if err != nil {
				return err
			}

			userApiService := user.NewUserAPIIntegrationService(user.UserAPIIntegrationServiceOptions{Ctx: cmd.Context()})

			validUser, err := userApiService.GetPlatformOwnerAndApplicationAdmin()
			if err != nil {
				return err
			}

			baseInfo := apiService.GetApplicationBaseInfoIDSlice(modelWrapper, validUser)

			_, err = apiService.ValidateApplication(
				baseInfo,
				_addApplicationOperation.applicationTechnology,
				_addApplicationOperation.applicationOwner,
				_addApplicationOperation.applicationCors,
			)
			if err != nil {
				return err
			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runAddCommand,
	}
	init.Flags().StringVarP(&_addApplicationOperation.applicationDescription, "application_description", "d", "", "enter application description")
	init.Flags().StringVarP(&_addApplicationOperation.applicationTechnology, "application_technology", "t", "", "enter application technology")
	init.Flags().StringVarP(&_addApplicationOperation.applicationOwner, "application_owner", "o", "", "enter application owner")
	init.Flags().StringVarP(&_addApplicationOperation.applicationCors, "application_cors", "c", "", "enter application cors")

	return init
}

func runAddCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	var applicationCors []contractApplication.Cors
	for _, cors := range strings.Split(_addApplicationOperation.applicationCors, ";") {
		applicationCors = append(applicationCors, contractApplication.Cors{
			Address: cors,
		})
	}

	applicationID, err := newApplicationAPIService(cmd.Context()).AddApplication(contractApplication.RequestApplication{
		Name:        _addApplicationOperation.applicationName,
		Description: _addApplicationOperation.applicationDescription,
		Technology:  _addApplicationOperation.applicationTechnology,
		Owner:       _addApplicationOperation.applicationOwner,
		Cors:        applicationCors,
	})

	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error creating application '%s': ", _addApplicationOperation.applicationName), err)
		return
	}

	common.CliRenderSuccessActionResourceOutput(width, applicationID, common.ResourceApplication, common.ActionCreated)
}
