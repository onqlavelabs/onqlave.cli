package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type UpdateClusterResponse struct {
	common.BaseErrorResponse
	Data contracts.ExistingClusterWithDetails `json:"data"`
}
