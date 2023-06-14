package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type GetClustersResponse struct {
	common.BaseErrorResponse
	Data contracts.GetArxsResponseWrapper `json:"data"`
}
