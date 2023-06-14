package requests

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
)

type AddClusterRequest struct {
	Cluster contracts.NewCluster `json:"cluster" validate:"required"`
}
