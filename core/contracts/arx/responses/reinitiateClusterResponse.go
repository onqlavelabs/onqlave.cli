package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type ReinitiateClusterResponse struct {
	common.BaseErrorResponse
	Data contracts.ClusterStatus `json:"data"`
}
