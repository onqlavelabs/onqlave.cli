package user

import (
	"context"
	"fmt"
	"github.com/onqlavelabs/onqlave.cli/internal/api"
	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/user"
)

type Service struct {
	opts ServiceOpt
}

type ServiceOpt struct {
	Ctx context.Context
}

func NewService(opts ServiceOpt) *Service {
	return &Service{opts: opts}
}

func (s *Service) GetUsers() (user.ListResponse, error) {
	tenantID := viper.GetString("tenant_id")
	userUrl := fmt.Sprintf("%s/%s/users", api.UrlBuilder(api.TenantName.String()), tenantID)

	response, err := api.Get[user.GetListResponse](userUrl)
	if err != nil {
		return user.ListResponse{}, err
	}

	return response.Data, nil
}

func (s *Service) GetPlatformOwnerAndClusterAdmin() (user.ListResponse, error) {
	tenantId := viper.Get("tenant_id")
	userUrl := fmt.Sprintf("%s/%s/users?roles=platform_owner,cluster_admin", api.UrlBuilder(api.TenantName.String()), tenantId)

	response, err := api.Get[user.GetListResponse](userUrl)
	if err != nil {
		return user.ListResponse{}, err
	}

	return response.Data, nil
}

func (s *Service) GetPlatformOwnerAndApplicationAdmin() (user.ListResponse, error) {
	tenantId := viper.Get("tenant_id")
	userUrl := fmt.Sprintf("%s/%s/users?roles=platform_owner,application_admin", api.UrlBuilder(api.TenantName.String()), tenantId)

	response, err := api.Get[user.GetListResponse](userUrl)
	if err != nil {
		return user.ListResponse{}, err
	}

	return response.Data, nil
}
