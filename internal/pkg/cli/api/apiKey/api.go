package apiKey

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/api_key"
	"github.com/onqlavelabs/onqlave.cli/core/enumerations"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/model"
)

type CommandOperation string

const (
	AddOperation    CommandOperation = "add"
	DeleteOperation CommandOperation = "delete"
)

var expectedOperationStatus = map[CommandOperation]enumerations.ApiKeyStatus{
	AddOperation:    enumerations.Active,
	DeleteOperation: enumerations.Deleted,
}

type ListKeysResponse struct {
	Keys []api_key.APIKey
}

type APIKeyIntegrationService struct {
	opts APIKeyIntegrationServiceOptions
}

type APIKeyIntegrationServiceOptions struct {
	Ctx context.Context
}

func NewAPIKeyIntegrationService(opts APIKeyIntegrationServiceOptions) *APIKeyIntegrationService {
	return &APIKeyIntegrationService{
		opts: opts,
	}
}

func (s *APIKeyIntegrationService) GetKeyBaseInfo() (api_key.APIKeyModels, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/keys/base", api.UrlBuilder(api.TenantName.String()), tenantId)

	response, err := api.Get[api_key.APIKeysResponse](clusterUrl)
	if err != nil {
		return api_key.APIKeyModels{}, model.NewAppError("GetKeyBaseInfo", "cli.server_error.key_base_info", nil, "get key base info failed", http.StatusInternalServerError).Wrap(err)
	}

	return response.Data.Model, nil
}

func (s *APIKeyIntegrationService) ValidateAPIKey(baseInfo api_key.APIKeyModels, appID, clusterID, appTech string) (bool, error) {
	var isClusterIDValid bool
	for _, cluster := range baseInfo.Arx {
		if cluster.ID == clusterID {
			isClusterIDValid = true
			break
		}
	}
	if !isClusterIDValid {
		return false, model.NewAppError("ValidateAPIKey", "cli.invalid.apikey_error", nil, "", http.StatusBadRequest).
			Wrap(errors.New("cluster id is invalid"))

	}

	var isAppIDValid bool
	var isTechValid bool
	for _, app := range baseInfo.Applications {
		if app.ID == appID && app.ApplicationTechnology.Id == appTech {
			isAppIDValid = true
			isTechValid = true
			break
		}
	}
	if !isAppIDValid || !isTechValid {
		return false, model.NewAppError("ValidateAPIKey", "cli.invalid.apikey_error", nil, "", http.StatusBadRequest).
			Wrap(errors.New("app id or app technology is invalid"))

	}

	return true, nil
}

func (s *APIKeyIntegrationService) CheckAPIKeyOperationStatus(keyId string, operation CommandOperation) (*api.APIIntegrationServiceOperationResult, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/keys/%s", api.UrlBuilder(api.TenantName.String()), tenantId, keyId)

	response, err := api.Get[api_key.APIKeyResponse](clusterUrl)
	message := "Checking api key operation status"
	if err != nil {
		return &api.APIIntegrationServiceOperationResult{Done: false, Result: message}, err
	}

	switch response.Data.Status {
	case enumerations.Failed.String():
		return &api.APIIntegrationServiceOperationResult{Done: false, Result: message}, fmt.Errorf("api key operation failed")
	case enumerations.Pending.String(), enumerations.Disabled.String():
		return &api.APIIntegrationServiceOperationResult{Done: false, Result: message}, nil
	case enumerations.Active.String(), enumerations.Deleted.String():
		if expectedOperationStatus[operation] == enumerations.ApiKeyStatus(response.Data.Status) {
			return &api.APIIntegrationServiceOperationResult{Done: true, Result: message}, nil
		}
		return &api.APIIntegrationServiceOperationResult{Done: false, Result: message}, nil
	default:
		return &api.APIIntegrationServiceOperationResult{Done: true, Result: message}, fmt.Errorf("provisioning state is invalid. please contact support. ")
	}
}

func (s *APIKeyIntegrationService) GetKeys() (ListKeysResponse, error) {
	tenantId := viper.Get("tenant_id")
	keyUrl := fmt.Sprintf("%s/%s/keys", api.UrlBuilder(api.TenantName.String()), tenantId)

	response, err := api.Get[api_key.APIKeysResponse](keyUrl)
	if err != nil {
		return ListKeysResponse{}, model.NewAppError("GetKeys", "cli.server_error.get_keys", nil, "get api keys failed", http.StatusInternalServerError).Wrap(err)
	}
	return ListKeysResponse{
		Keys: response.Data.APIKeys,
	}, nil
}

func (s *APIKeyIntegrationService) DeleteKey(keyId string) (string, error) {
	tenantId := viper.Get("tenant_id")
	keyUrl := fmt.Sprintf("%s/%s/keys/%s", api.UrlBuilder(api.TenantName.String()), tenantId, keyId)

	_, err := api.Delete[map[string]interface{}](keyUrl)
	if err != nil {
		return "", model.NewAppError("DeleteKey", "cli.server_error.delete_key", nil, "delete api key failed", http.StatusInternalServerError).Wrap(err)
	}

	return keyId, nil
}

func (s *APIKeyIntegrationService) AddKey(contract api_key.CreateAPIKey) (string, error) {
	tenantId := viper.Get("tenant_id")
	keyUrl := fmt.Sprintf("%s/%s/keys", api.UrlBuilder(api.TenantName.String()), tenantId)

	request := api_key.CreateAPIKeyRequest{
		APIKey: contract,
	}
	response, err := api.Post[api_key.APIKeyResponse](keyUrl, request)
	if err != nil {
		return "", model.NewAppError("AddKey", "cli.server_error.add_key", nil, "add api key failed", http.StatusInternalServerError).Wrap(err)
	}

	return response.Data.ID, nil
}

func (s *APIKeyIntegrationService) GetKeyDetail(keyID string) (api_key.APIKey, error) {
	tenantId := viper.Get("tenant_id")
	keyUrl := fmt.Sprintf("%s/%s/keys/%s", api.UrlBuilder(api.TenantName.String()), tenantId, keyID)

	response, err := api.Get[api_key.APIKeyResponse](keyUrl)
	if err != nil {
		return api_key.APIKey{}, model.NewAppError("GetKeyDetail", "cli.server_error.get_key_detail", nil, "get api key detail failed", http.StatusInternalServerError).Wrap(err)
	}
	return response.Data, nil
}
