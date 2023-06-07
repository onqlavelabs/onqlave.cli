package requests

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"

type UpdateClusterRequest struct {
	Cluster contracts.UpdateCluster `json:"cluster" validate:"required"`
}
