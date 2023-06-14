package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/application"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type GetApplicationsResponse struct {
	common.BaseErrorResponse
	Data contracts.GetApplicationsResponseWrapper `json:"data"`
}
