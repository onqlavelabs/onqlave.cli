package arx

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/internal/app/tenant/cluster/enums"
	acl "github.com/onqlavelabs/onqlave.cli/internal/pkg/acl/contracts"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/model"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts/requests"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts/responses"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/utils"
)

type CommandOperation string

const (
	UpdateOperation CommandOperation = "update"
	RetryOperation  CommandOperation = "retry"
	AddOperation    CommandOperation = "add"
	UnsealOperation CommandOperation = "unseal"
	SealOperation   CommandOperation = "seal"
	DeleteOperation CommandOperation = "delete"
)

var expectedOperationStatus = map[CommandOperation]enums.ClusterStatus{
	UpdateOperation: enums.ClusterActive,
	RetryOperation:  enums.ClusterActive,
	AddOperation:    enums.ClusterActive,
	UnsealOperation: enums.ClusterActive,
	SealOperation:   enums.ClusterSealed,
	DeleteOperation: enums.ClusterDeleted,
}

type ArxBaseInfo struct {
	ProviderIDs          []string
	PlanIDs              []string
	PurposeIDs           []string
	EncryptionMethodIDs  []string
	RotationCycleIDs     []string
	CloudProviderRegions map[string][]string
}

type GetDetailArxResponse struct {
	ID  string `json:"id"`
	Acl acl.ACL
	contracts.ClusterDetail
}

type ListArxResponse struct {
	Clusters []contracts.ExistingClusterWithDetails `json:"clusters"`
}

type ArxAPIIntegrationService struct {
	opts ArxAPIIntegrationServiceOptions
}

type ArxAPIIntegrationServiceOptions struct {
	Ctx context.Context
}

func NewArxAPIIntegrationService(opts ArxAPIIntegrationServiceOptions) *ArxAPIIntegrationService {
	return &ArxAPIIntegrationService{
		opts: opts,
	}
}

func (s *ArxAPIIntegrationService) CheckArxOperationState(clusterId string, operation CommandOperation) (*api.APIIntegrationServiceOperationResult, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters/%s/state", api.UrlBuilder(api.TenantName.String()), tenantId, clusterId)

	response, err := api.Get[responses.GetClusterStateResponse](clusterUrl)
	message := "Waiting for arx operation to complete."
	if err != nil {
		return &api.APIIntegrationServiceOperationResult{Done: false, Result: message}, err
	}

	switch response.Data.State {
	case enums.ClusterFailed.String():
		return &api.APIIntegrationServiceOperationResult{Done: false, Result: message}, fmt.Errorf(response.Data.Message)
	case enums.ClusterInactive.String(),
		enums.ClusterPending.String(),
		enums.ClusterInitiated.String(),
		enums.ClusterReInitiated.String(),
		enums.ClusterUnsealed.String():
		return &api.APIIntegrationServiceOperationResult{Done: false, Result: message}, nil
	case enums.ClusterActive.String(),
		enums.ClusterSealed.String(),
		enums.ClusterDeleted.String():
		if expectedOperationStatus[operation] == enums.ClusterStatus(response.Data.State) {
			return &api.APIIntegrationServiceOperationResult{Done: true, Result: message}, nil
		}
		return &api.APIIntegrationServiceOperationResult{Done: false, Result: message}, nil
	default:
		return &api.APIIntegrationServiceOperationResult{Done: true, Result: message}, fmt.Errorf("provisioning state is invalid. please contact support. ")
	}
}

func (s *ArxAPIIntegrationService) GetArxBaseInfo() (contracts.ArxModelWrapper, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters/base", api.UrlBuilder(api.TenantName.String()), tenantId)

	response, err := api.Get[responses.GetClusterBaseInfoWrapper](clusterUrl)
	if err != nil {
		return contracts.ArxModelWrapper{}, model.NewAppError("GetClusterBaseInfo", "cli.server_error.cluster_base_info", nil, "get cluster base info failed", http.StatusInternalServerError).Wrap(err)
	}

	return response.Data, nil
}

func (s *ArxAPIIntegrationService) GetArxDetail(clusterID string) (*GetDetailArxResponse, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters/%s/state?detail=true", api.UrlBuilder(api.TenantName.String()), tenantId, clusterID)

	response, err := api.Get[responses.GetClusterStateResponse](clusterUrl)
	if err != nil {
		return nil, model.NewAppError("GetClusterDetail", "cli.server_error.cluster_detail", nil, "get cluster detail failed", http.StatusInternalServerError).Wrap(err)
	}

	return &GetDetailArxResponse{
		string(response.Data.ID),
		response.Data.ACL,
		*response.Data.Cluster,
	}, nil
}

func (s *ArxAPIIntegrationService) GetArxBaseInfoIDSlice(data contracts.ArxModelWrapper) ArxBaseInfo {
	var baseInfo ArxBaseInfo
	var cloudProviderRegions = make(map[string][]string)
	for _, provider := range data.Providers {
		if !utils.BoolValue(provider.Enable) {
			continue
		}
		baseInfo.ProviderIDs = append(baseInfo.ProviderIDs, provider.ID)
		if provider.Regions != nil {
			for _, region := range provider.Regions {
				if utils.BoolValue(region.Enable) {
					cloudProviderRegions[provider.ID] = append(cloudProviderRegions[provider.ID], region.ID)
				}
			}
		}
	}
	baseInfo.CloudProviderRegions = cloudProviderRegions

	for _, plan := range data.Plans {
		if utils.BoolValue(plan.Enable) {
			baseInfo.PlanIDs = append(baseInfo.PlanIDs, plan.ID)
		}
	}

	for _, purpose := range data.Purposes {
		if utils.BoolValue(purpose.Enable) {
			baseInfo.PurposeIDs = append(baseInfo.PurposeIDs, purpose.ID)
		}
	}

	for _, encryptionMethod := range data.EncryptionMethods {
		if utils.BoolValue(encryptionMethod.Enable) {
			baseInfo.EncryptionMethodIDs = append(baseInfo.EncryptionMethodIDs, encryptionMethod.ID)
		}
	}

	for _, rotationCycle := range data.RotationCycles {
		if utils.BoolValue(rotationCycle.Enable) {
			baseInfo.RotationCycleIDs = append(baseInfo.RotationCycleIDs, rotationCycle.ID)
		}
	}

	return baseInfo
}

func (s *ArxAPIIntegrationService) ValidateArx(
	baseInfo ArxBaseInfo,
	clusterProvider string,
	clusterType string,
	clusterPurpose string,
	clusterRegion string,
	encryptionMethod string,
	rotationCycle string,
) (bool, error) {
	if !utils.Contains(baseInfo.ProviderIDs, clusterProvider) {
		return false, model.NewAppError("ValidateCluster", "cli.invalid.cluster_cloud_provider", nil, "", http.StatusBadRequest).
			Wrap(fmt.Errorf("invalid cluster provider - must be in (%v)", strings.Join(baseInfo.ProviderIDs, ", ")))
	}

	if regions, ok := baseInfo.CloudProviderRegions[clusterProvider]; ok {
		if !utils.Contains(regions, clusterRegion) {
			return false, model.NewAppError("ValidateCluster", "cli.invalid.cluster_cloud_provider_region", nil, "", http.StatusBadRequest).
				Wrap(fmt.Errorf("invalid cluster provider - must be in (%v)", strings.Join(regions, ", ")))
		}
	}

	if !utils.Contains(baseInfo.PlanIDs, clusterType) {
		return false, model.NewAppError("ValidateCluster", "cli.invalid.cluster_plan", nil, "", http.StatusBadRequest).
			Wrap(fmt.Errorf("invalid cluster type - must be in (%v)", strings.Join(baseInfo.PlanIDs, ", ")))
	}

	if !utils.Contains(baseInfo.PurposeIDs, clusterPurpose) {
		return false, model.NewAppError("ValidateCluster", "cli.invalid.cluster_purpose", nil, "", http.StatusBadRequest).
			Wrap(fmt.Errorf("invalid cluster purpose - must be in (%v)", strings.Join(baseInfo.PurposeIDs, ", ")))
	}

	if !utils.Contains(baseInfo.EncryptionMethodIDs, encryptionMethod) {
		return false, model.NewAppError("ValidateCluster", "cli.invalid.cluster_encryption_method", nil, "", http.StatusBadRequest).
			Wrap(fmt.Errorf("invalid cluster encryption method - must be in (%v)", strings.Join(baseInfo.EncryptionMethodIDs, ", ")))
	}

	if !utils.Contains(baseInfo.RotationCycleIDs, rotationCycle) {
		return false, model.NewAppError("ValidateCluster", "cli.invalid.cluster_rotation_cycle", nil, "", http.StatusBadRequest).
			Wrap(fmt.Errorf("invalid cluster rotation cycle - must be in (%v)", strings.Join(baseInfo.RotationCycleIDs, ", ")))
	}

	return true, nil
}

func (s ArxAPIIntegrationService) ValidateEditArxRequest(
	baseInfo ArxBaseInfo,
	clusterProvider string,
	clusterRegion string,
	rotationCycle string,
) (bool, error) {
	if regions, ok := baseInfo.CloudProviderRegions[clusterProvider]; ok {
		if !utils.Contains(regions, clusterRegion) {
			return false, model.NewAppError("ValidateCluster", "cli.invalid.cluster_cloud_provider_region", nil, "", http.StatusBadRequest).
				Wrap(fmt.Errorf("invalid cluster provider - must be in (%v)", strings.Join(regions, ", ")))
		}
	}

	if !utils.Contains(baseInfo.RotationCycleIDs, rotationCycle) {
		return false, model.NewAppError("ValidateCluster", "cli.invalid.cluster_rotation_cycle", nil, "", http.StatusBadRequest).
			Wrap(fmt.Errorf("invalid cluster rotation cycle - must be in (%v)", strings.Join(baseInfo.RotationCycleIDs, ", ")))
	}

	return true, nil
}

func (s *ArxAPIIntegrationService) AddArx(addClusterRequest contracts.NewCluster) (string, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters", api.UrlBuilder(api.TenantName.String()), tenantId)

	request := requests.AddClusterRequest{
		Cluster: addClusterRequest,
	}
	response, err := api.Post[responses.AddClusterResponse](clusterUrl, request)
	if err != nil {
		return "", model.NewAppError("AddCluster", "cli.server_error.cluster", nil, "create cluster failed", http.StatusInternalServerError).Wrap(err)
	}
	return string(response.Data.ID), nil
}

func (s *ArxAPIIntegrationService) RetryAddArx(clusterId string) (string, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters/%s/retry", api.UrlBuilder(api.TenantName.String()), tenantId, clusterId)

	request := struct{}{}
	response, err := api.Post[responses.ReinitiateClusterResponse](clusterUrl, request)
	if err != nil {
		return "", err
	}
	return string(response.Data.ID), nil
}

func (s *ArxAPIIntegrationService) DeleteArx(clusterId string) (string, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters/%s", api.UrlBuilder(api.TenantName.String()), tenantId, clusterId)

	response, err := api.Delete[responses.DeleteClusterResponse](clusterUrl)
	if err != nil {
		return "", model.NewAppError("DeleteCluster", "cli.server_error.delete_cluster", nil, "delete cluster failed", http.StatusInternalServerError).Wrap(err)
	}
	return string(response.Data.ID), nil
}

func (s *ArxAPIIntegrationService) SealArx(clusterId string) (string, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters/%s/seal", api.UrlBuilder(api.TenantName.String()), tenantId, clusterId)

	request := struct{}{}
	response, err := api.Post[responses.SealClusterResponse](clusterUrl, request)
	if err != nil {
		return "", model.NewAppError("SealCluster", "cli.server_error.seal_cluster", nil, "seal cluster failed", http.StatusInternalServerError).Wrap(err)
	}
	return string(response.Data.ID), nil
}

func (s *ArxAPIIntegrationService) UnsealArx(clusterId string) (string, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters/%s/unseal", api.UrlBuilder(api.TenantName.String()), tenantId, clusterId)

	request := struct{}{}
	response, err := api.Post[responses.UnsealClusterResponse](clusterUrl, request)
	if err != nil {
		return "", model.NewAppError("UnsealCluster", "cli.server_error.unseal_cluster", nil, "unseal cluster failed", http.StatusInternalServerError).Wrap(err)
	}
	return string(response.Data.ID), nil
}

func (s *ArxAPIIntegrationService) SetDefaultArx(clusterId string) (string, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters/%s/default", api.UrlBuilder(api.TenantName.String()), tenantId, clusterId)

	request := struct{}{}
	response, err := api.Put[responses.UpdateDefaultClusterResponse](clusterUrl, request)
	if err != nil {
		return "", model.NewAppError("SetDefaultCluster", "cli.server_error.set_default_cluster", nil, "set default cluster failed", http.StatusInternalServerError).Wrap(err)
	}
	return string(response.Data.ID), nil
}

func (s *ArxAPIIntegrationService) UpdateArx(contract contracts.UpdateCluster) (string, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters/%s", api.UrlBuilder(api.TenantName.String()), tenantId, contract.ID)
	request := requests.UpdateClusterRequest{
		Cluster: contract,
	}
	response, err := api.Put[responses.UpdateClusterResponse](clusterUrl, request)
	if err != nil {
		return "", model.NewAppError("UpdateCluster", "cli.server_error.update_cluster", nil, "update cluster failed", http.StatusInternalServerError).Wrap(err)
	}
	return string(response.Data.ID), nil
}

func (s *ArxAPIIntegrationService) GetArx() (ListArxResponse, error) {
	tenantId := viper.Get("tenant_id")
	clusterUrl := fmt.Sprintf("%s/%s/clusters", api.UrlBuilder(api.TenantName.String()), tenantId)

	response, err := api.Get[responses.GetClustersResponse](clusterUrl)
	if err != nil {
		return ListArxResponse{}, model.NewAppError("GetClusters", "cli.server_error.get_clusters", nil, "get clusters failed", http.StatusInternalServerError).Wrap(err)
	}
	return ListArxResponse{
		Clusters: response.Data.Clusters,
	}, nil
}