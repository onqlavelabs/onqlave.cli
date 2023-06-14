package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/application"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type SealApplicationResponse struct {
	common.BaseErrorResponse
	Data contracts.ApplicationStatus `json:"data"`
}
