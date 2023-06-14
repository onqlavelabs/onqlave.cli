package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type AddClusterResponse struct {
	common.BaseErrorResponse
	Data contracts.ExistingClusterWithDetails `json:"data"`
}
