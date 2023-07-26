package application

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api/application"
	"github.com/onqlavelabs/onqlave.cli/internal/api/user"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	contractApplication "github.com/onqlavelabs/onqlave.core/contracts/application"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type editApplicationOperation struct {
	applicationID          string
	applicationName        string
	applicationDescription string
	applicationTechnology  string
	applicationOwner       string
	applicationCors        string
	start                  time.Time
}

var _editApplicationOperation editApplicationOperation

func updateCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "update",
		Short:   "update application by ID and attributes",
		Long:    "This command is used to update application by ID. Application ID is required.",
		Example: "onqlave application update",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("ApplicationID is required")))
			}
			_editApplicationOperation.applicationID = args[0]
			return nil
		},
		// used to overwrite/skip the parent commands persistentPreRunE func
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if viper.GetBool(common.FlagDebug) {
				_editApplicationOperation.start = time.Now()
			}

			apiService := application.NewService(application.ServiceOpt{Ctx: cmd.Context()})
			applicationDetail, err := apiService.GetApplication(_editApplicationOperation.applicationID)
			if err != nil {
				return fmt.Errorf("There was an error describing application '%s': %s", _describeApplication.applicationId, err)
			}

			if _editApplicationOperation.applicationName == "" {
				_editApplicationOperation.applicationName = applicationDetail.Name
			}

			if _editApplicationOperation.applicationTechnology == "" {
				_editApplicationOperation.applicationTechnology = applicationDetail.Technology
			}

			if _editApplicationOperation.applicationOwner == "" {
				_editApplicationOperation.applicationOwner = applicationDetail.Owner
			}

			if _editApplicationOperation.applicationCors == "" {
				for _, corsAddress := range applicationDetail.Cors {
					_editApplicationOperation.applicationCors = fmt.Sprintf("%s;%s", _editApplicationOperation.applicationCors, corsAddress.Address)
				}
				_editApplicationOperation.applicationCors = strings.TrimLeft(_editApplicationOperation.applicationCors, ";")
			}

			modelWrapper, err := apiService.GetBaseApplication()
			if err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}

			userApiService := user.NewService(user.ServiceOpt{Ctx: cmd.Context()})
			validUser, err := userApiService.GetPlatformOwnerAndApplicationAdmin()
			if err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}

			baseInfo := apiService.GetApplicationBaseInfoIDSlice(modelWrapper, validUser)

			_, err = apiService.ValidateApplication(
				baseInfo,
				_editApplicationOperation.applicationTechnology,
				_editApplicationOperation.applicationOwner,
				_editApplicationOperation.applicationCors,
			)
			if err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runEditCommand,
	}
	init.Flags().StringVarP(&_editApplicationOperation.applicationName, "application_name", "n", "", "enter application name")
	init.Flags().StringVarP(&_editApplicationOperation.applicationDescription, "application_description", "d", "", "enter application description")
	init.Flags().StringVarP(&_editApplicationOperation.applicationTechnology, "application_technology", "t", "", "enter application technology")
	init.Flags().StringVarP(&_editApplicationOperation.applicationOwner, "application_owner", "o", "", "enter application owner")
	init.Flags().StringVarP(&_editApplicationOperation.applicationCors, "application_cors", "c", "", "enter application cors")

	return init
}

func runEditCommand(cmd *cobra.Command, args []string) {
	start := time.Now()
	defer common.LogResponseTime(start)

	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	var applicationCors []contractApplication.Cors
	for _, cors := range strings.Split(_editApplicationOperation.applicationCors, ";") {
		applicationCors = append(applicationCors, contractApplication.Cors{
			Address: cors,
		})
	}

	applicationID, err := newApplicationAPIService(cmd.Context()).EditApplication(
		_editApplicationOperation.applicationID,
		contractApplication.RequestApplication{
			Name:        _editApplicationOperation.applicationName,
			Description: _editApplicationOperation.applicationDescription,
			Technology:  _editApplicationOperation.applicationTechnology,
			Owner:       _editApplicationOperation.applicationOwner,
			Cors:        applicationCors,
		})

	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("There was an error updating application '%s': ", _addApplicationOperation.applicationName), err)
		return
	}

	common.CliRenderSuccessActionResourceOutput(width, applicationID, common.ResourceApplication, common.ActionUpdated)
}
