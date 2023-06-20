package user

import (
	"context"
	"fmt"
	api2 "github.com/onqlavelabs/onqlave.cli/internal/cli/api"

	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/user"
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

func (s *UserAPIIntegrationService) GetUsers() (user.ListResponse, error) {
	tenantID := viper.GetString("tenant_id")
	userUrl := fmt.Sprintf("%s/%s/users", api2.UrlBuilder(api2.TenantName.String()), tenantID)

	response, err := api2.Get[user.GetListResponse](userUrl)
	if err != nil {
		return user.ListResponse{}, err
	}

	return response.Data, nil
}

func (s *UserAPIIntegrationService) GetPlatformOwnerAndClusterAdmin() (user.ListResponse, error) {
	tenantId := viper.Get("tenant_id")
	userUrl := fmt.Sprintf("%s/%s/users?roles=platform_owner,cluster_admin", api2.UrlBuilder(api2.TenantName.String()), tenantId)

	response, err := api2.Get[user.GetListResponse](userUrl)
	if err != nil {
		return user.ListResponse{}, err
	}

	return response.Data, nil
}

func (s *UserAPIIntegrationService) GetPlatformOwnerAndApplicationAdmin() (user.ListResponse, error) {
	tenantId := viper.Get("tenant_id")
	userUrl := fmt.Sprintf("%s/%s/users?roles=platform_owner,application_admin", api2.UrlBuilder(api2.TenantName.String()), tenantId)

	response, err := api2.Get[user.GetListResponse](userUrl)
	if err != nil {
		return user.ListResponse{}, err
	}

	return response.Data, nil
}
