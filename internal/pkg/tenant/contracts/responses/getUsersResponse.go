package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type GetUsersResponse struct {
	common.BaseErrorResponse
	Data contracts.GetUsersResponse `json:"data"`
}
