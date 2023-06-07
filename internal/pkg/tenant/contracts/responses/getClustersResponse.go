package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type GetClustersResponse struct {
	common.BaseErrorResponse
	Data contracts.GetArxsResponseWrapper `json:"data"`
}
