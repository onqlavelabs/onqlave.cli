package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type DeleteClusterResponse struct {
	common.BaseErrorResponse
	Data contracts.ClusterStatus `json:"data"`
}
