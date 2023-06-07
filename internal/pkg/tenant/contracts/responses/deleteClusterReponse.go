package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type DeleteClusterResponse struct {
	common.BaseErrorResponse
	Data contracts.ClusterStatus `json:"data"`
}
