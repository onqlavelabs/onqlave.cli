package user

import (
	"context"
	"fmt"

	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts/responses"
)

type UserAPIIntegrationService struct {
	opts UserAPIIntegrationServiceOptions
}

type UserAPIIntegrationServiceOptions struct {
	Ctx context.Context
}

func NewUserAPIIntegrationService(opts UserAPIIntegrationServiceOptions) *UserAPIIntegrationService {
	return &UserAPIIntegrationService{
		opts: opts,
	}
}

func (s *UserAPIIntegrationService) GetUsers() (contracts.GetUsersResponse, error) {
	tenantID := viper.GetString("tenant_id")
	userUrl := fmt.Sprintf("%s/%s/users", api.UrlBuilder(api.TenantName.String()), tenantID)

	response, err := api.Get[responses.GetUsersResponse](userUrl)
	if err != nil {
		return contracts.GetUsersResponse{}, err
	}

	return response.Data, nil
}

func (s *UserAPIIntegrationService) GetPlatformOwnerAndClusterAdmin() (contracts.GetUsersResponse, error) {
	tenantId := viper.Get("tenant_id")
	userUrl := fmt.Sprintf("%s/%s/users?roles=platform_owner,cluster_admin", api.UrlBuilder(api.TenantName.String()), tenantId)

	response, err := api.Get[responses.GetUsersResponse](userUrl)
	if err != nil {
		return contracts.GetUsersResponse{}, err
	}

	return response.Data, nil
}

func (s *UserAPIIntegrationService) GetPlatformOwnerAndApplicationAdmin() (contracts.GetUsersResponse, error) {
	tenantId := viper.Get("tenant_id")
	userUrl := fmt.Sprintf("%s/%s/users?roles=platform_owner,application_admin", api.UrlBuilder(api.TenantName.String()), tenantId)

	response, err := api.Get[responses.GetUsersResponse](userUrl)
	if err != nil {
		return contracts.GetUsersResponse{}, err
	}

	return response.Data, nil
}
