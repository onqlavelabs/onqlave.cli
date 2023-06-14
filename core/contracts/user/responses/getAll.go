package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type GetUsersResponse struct {
	common.BaseErrorResponse
	Data contracts.GetUsersResponse `json:"data"`
}
