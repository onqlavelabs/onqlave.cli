package requests

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"

type AddClusterRequest struct {
	Cluster contracts.NewCluster `json:"cluster" validate:"required"`
}
