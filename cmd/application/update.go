package application

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
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
}

var updateApplication editApplicationOperation

func updateCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "update",
		Short:   "update application by ID and attributes",
		Long:    "This command is used to update application by ID. Application ID is required.",
		Example: "onqlave application update",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.CliRenderErr(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("ApplicationID is required")))
			}
			updateApplication.applicationID = args[0]
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			apiService := application.NewService(application.ServiceOpt{Ctx: cmd.Context()})
			applicationDetail, err := apiService.GetApplication(updateApplication.applicationID)
			if err != nil {
				return fmt.Errorf("there was an error describing application '%s': %s", updateApplication.applicationID, err)
			}

			if updateApplication.applicationName == "" {
				updateApplication.applicationName = applicationDetail.Name
			}

			if updateApplication.applicationTechnology == "" {
				updateApplication.applicationTechnology = applicationDetail.Technology
			}

			if updateApplication.applicationOwner == "" {
				updateApplication.applicationOwner = applicationDetail.Owner
			}

			if updateApplication.applicationCors == "" {
				for _, corsAddress := range applicationDetail.Cors {
					updateApplication.applicationCors = fmt.Sprintf("%s;%s", updateApplication.applicationCors, corsAddress.Address)
				}
				updateApplication.applicationCors = strings.TrimLeft(updateApplication.applicationCors, ";")
			}

			modelWrapper, err := apiService.GetBaseApplication()
			if err != nil {
				return common.CliRenderErr(cmd, err)
			}

			userApiService := user.NewService(user.ServiceOpt{Ctx: cmd.Context()})
			validUser, err := userApiService.GetPlatformOwnerAndApplicationAdmin()
			if err != nil {
				return common.CliRenderErr(cmd, err)
			}

			baseInfo := apiService.GetApplicationBaseInfoIDSlice(modelWrapper, validUser)

			_, err = apiService.ValidateApplication(
				baseInfo,
				updateApplication.applicationTechnology,
				updateApplication.applicationOwner,
				updateApplication.applicationCors,
			)
			if err != nil {
				return common.CliRenderErr(cmd, err)
			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runEditCommand,
	}
	init.Flags().StringVarP(&updateApplication.applicationName, "application_name", "n", "", "enter application name")
	init.Flags().StringVarP(&updateApplication.applicationDescription, "application_description", "d", "", "enter application description")
	init.Flags().StringVarP(&updateApplication.applicationTechnology, "application_technology", "t", "", "enter application technology")
	init.Flags().StringVarP(&updateApplication.applicationOwner, "application_owner", "o", "", "enter application owner")
	init.Flags().StringVarP(&updateApplication.applicationCors, "application_cors", "c", "", "enter application cors")

	return init
}

func runEditCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	var applicationCors []contractApplication.Cors
	for _, cors := range strings.Split(updateApplication.applicationCors, ";") {
		applicationCors = append(applicationCors, contractApplication.Cors{
			Address: cors,
		})
	}

	applicationID, err := newApplicationAPIService(cmd.Context()).EditApplication(
		updateApplication.applicationID,
		contractApplication.RequestApplication{
			Name:        updateApplication.applicationName,
			Description: updateApplication.applicationDescription,
			Technology:  updateApplication.applicationTechnology,
			Owner:       updateApplication.applicationOwner,
			Cors:        applicationCors,
		})

	if err != nil {
		common.RenderCLIOutputError(fmt.Sprintf("there was an error updating application '%s': ", updateApplication.applicationName), err)
		return
	}

	common.CliRenderSuccessActionResourceOutput(width, applicationID, common.ResourceApplication, common.ActionUpdated)
}
