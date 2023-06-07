package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type GetApplicationsResponse struct {
	common.BaseErrorResponse
	Data contracts.GetApplicationsResponseWrapper `json:"data"`
}
