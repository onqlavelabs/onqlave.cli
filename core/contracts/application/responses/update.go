package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/application"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type UpdateApplicationResponse struct {
	common.BaseErrorResponse
	Data contracts.ExistingApplicationWithDetails `json:"data"`
}
