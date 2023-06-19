package application

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/core/contracts"
	contractsApp "github.com/onqlavelabs/onqlave.cli/core/contracts/application"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/model"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/utils"
)

type CommandOperation string

type ApplicationBaseInfo struct {
	Technologies map[string]bool
	User         []string
}

type ApplicationAPIIntegrationService struct {
	opts ApplicationAPIIntegrationServiceOptions
}

type ApplicationAPIIntegrationServiceOptions struct {
	Ctx context.Context
}

func NewApplicationAPIIntegrationService(opts ApplicationAPIIntegrationServiceOptions) *ApplicationAPIIntegrationService {
	return &ApplicationAPIIntegrationService{
		opts: opts,
	}
}

func (s *ApplicationAPIIntegrationService) ValidateApplication(baseInfo ApplicationBaseInfo, technology, owner, corsIp string) (bool, error) {
	if !utils.Contains(baseInfo.User, owner) {
		return false, model.NewAppError("ValidateApplication", "cli.invalid.application_owner", nil, "", http.StatusBadRequest).
			Wrap(fmt.Errorf("invalid owner - must be in (%v)", strings.Join(baseInfo.User, ", ")))
	}

	if corsRequired, ok := baseInfo.Technologies[technology]; ok {
		if corsRequired && corsIp == "" {
			return false, model.NewAppError("ValidateApplication", "cli.invalid.application_cors", nil, "", http.StatusBadRequest).
				Wrap(fmt.Errorf("invalid cors: technology selected requires cors"))
		}
		return true, nil
	}

	validTechnologies := ""
	for tech := range baseInfo.Technologies {
		validTechnologies = fmt.Sprintf("%s,%s", validTechnologies, tech)
	}

	return false, model.NewAppError("ValidateApplication", "cli.invalid.application_technology", nil, "", http.StatusBadRequest).
		Wrap(fmt.Errorf("invalid technology - must be in (%v)", strings.TrimLeft(validTechnologies, ",")))
}

func (s *ApplicationAPIIntegrationService) GetApplicationBaseInfoIDSlice(modelWrapper contractsApp.Technologies, validUser contracts.GetUsersResponse) ApplicationBaseInfo {
	baseInfo := ApplicationBaseInfo{
		Technologies: map[string]bool{},
		User:         []string{},
	}

	for _, technology := range modelWrapper.Technologies {
		baseInfo.Technologies[technology.Id] = technology.Cors
	}

	for _, user := range validUser.Users {
		baseInfo.User = append(baseInfo.User, user.ID)
	}

	return baseInfo
}

func (s *ApplicationAPIIntegrationService) GetBaseApplication() (contractsApp.Technologies, error) {
	tenantId := viper.Get("tenant_id")
	applicationUrl := fmt.Sprintf("%s/%s/applications/base", api.UrlBuilder(api.TenantName.String()), tenantId)

	response, err := api.Get[contractsApp.BaseResponse](applicationUrl)
	if err != nil {
		return contractsApp.Technologies{}, model.NewAppError("Get Base Application", "cli.server_error.application", nil, "get base application failed", http.StatusInternalServerError).Wrap(err)
	}
	return response.Data, nil
}

func (s *ApplicationAPIIntegrationService) AddApplication(addApplicationRequest contractsApp.Application) (string, error) {
	tenantId := viper.Get("tenant_id")
	applicationUrl := fmt.Sprintf("%s/%s/applications", api.UrlBuilder(api.TenantName.String()), tenantId)

	request := contractsApp.Request{
		Application: addApplicationRequest,
	}
	response, err := api.Post[contractsApp.DetailResponse](applicationUrl, request)
	if err != nil {
		return "", model.NewAppError("CreateApplication", "cli.server_error.create_application", nil, "create application failed", http.StatusInternalServerError).Wrap(err)
	}
	return string(response.Data.ID), nil
}

func (s *ApplicationAPIIntegrationService) EditApplication(applicationID string, editApplicationRequest contractsApp.Application) (string, error) {
	tenantId := viper.Get("tenant_id")
	applicationUrl := fmt.Sprintf("%s/%s/applications/%s", api.UrlBuilder(api.TenantName.String()), tenantId, applicationID)

	request := contractsApp.Request{
		Application: editApplicationRequest,
	}
	response, err := api.Put[contractsApp.DetailResponse](applicationUrl, request)
	if err != nil {
		return "", model.NewAppError("EditApplication", "cli.server_error.edit_application", nil, "create application failed", http.StatusInternalServerError).Wrap(err)
	}
	return string(response.Data.ID), nil
}

func (s *ApplicationAPIIntegrationService) GetApplication(applicationID string) (contractsApp.Detail, error) {
	tenantId := viper.Get("tenant_id")
	applicationUrl := fmt.Sprintf("%s/%s/applications/%s", api.UrlBuilder(api.TenantName.String()), tenantId, applicationID)

	response, err := api.Get[contractsApp.DetailResponse](applicationUrl)
	if err != nil {
		return contractsApp.Detail{}, model.NewAppError("DescribeApplication", "cli.server_error.describe_applications", nil, "describe application failed", http.StatusInternalServerError).Wrap(err)
	}
	return response.Data, nil
}

func (s *ApplicationAPIIntegrationService) GetApplications() ([]contractsApp.Detail, error) {
	tenantId := viper.Get("tenant_id")
	applicationUrl := fmt.Sprintf("%s/%s/applications", api.UrlBuilder(api.TenantName.String()), tenantId)

	response, err := api.Get[contractsApp.ListResponse](applicationUrl)
	if err != nil {
		return []contractsApp.Detail{}, model.NewAppError("GetApplications", "cli.server_error.applications", nil, "get application failed", http.StatusInternalServerError).Wrap(err)
	}

	return response.Data.Applications, nil
}

func (s *ApplicationAPIIntegrationService) ArchiveApplication(applicationID string) error {
	tenantId := viper.Get("tenant_id")
	applicationUrl := fmt.Sprintf("%s/%s/applications/%s/archive", api.UrlBuilder(api.TenantName.String()), tenantId, applicationID)
	request := struct{}{}
	_, err := api.Post[contractsApp.DetailResponse](applicationUrl, request)
	if err != nil {
		return model.NewAppError("ArchiveApplications", "cli.server_error.archive_applications", nil, "archive application failed", http.StatusInternalServerError).Wrap(err)
	}
	return nil
}

func (s *ApplicationAPIIntegrationService) DisableApplication(applicationID string) error {
	tenantId := viper.Get("tenant_id")
	applicationUrl := fmt.Sprintf("%s/%s/applications/%s/disable", api.UrlBuilder(api.TenantName.String()), tenantId, applicationID)
	request := struct{}{}
	_, err := api.Post[contractsApp.DetailResponse](applicationUrl, request)
	if err != nil {
		return model.NewAppError("DisableApplications", "cli.server_error.disable_applications", nil, "disable application failed", http.StatusInternalServerError).Wrap(err)
	}
	return nil
}

func (s *ApplicationAPIIntegrationService) EnableApplication(applicationID string) error {
	tenantId := viper.Get("tenant_id")
	applicationUrl := fmt.Sprintf("%s/%s/applications/%s/enable", api.UrlBuilder(api.TenantName.String()), tenantId, applicationID)
	request := struct{}{}
	_, err := api.Post[contractsApp.DetailResponse](applicationUrl, request)
	if err != nil {
		return model.NewAppError("EnableApplications", "cli.server_error.enable_applications", nil, "enable application failed", http.StatusInternalServerError).Wrap(err)
	}
	return nil
}
