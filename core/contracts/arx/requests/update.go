package requests

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
)

type UpdateClusterRequest struct {
	Cluster contracts.UpdateCluster `json:"cluster" validate:"required"`
}
