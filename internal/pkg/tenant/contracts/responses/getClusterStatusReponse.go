package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type GetClusterStateResponse struct {
	common.BaseErrorResponse
	Data contracts.ClusterStatus `json:"data"`
}
