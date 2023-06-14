package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type UnsealClusterResponse struct {
	common.BaseErrorResponse
	Data contracts.ClusterStatus `json:"data"`
}
