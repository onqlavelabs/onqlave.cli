package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type GetClusterStateResponse struct {
	common.BaseErrorResponse
	Data contracts.ClusterStatus `json:"data"`
}
