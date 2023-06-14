package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/api_key"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type AddAPIKeyResponse struct {
	common.BaseErrorResponse
	Data contracts.APIKey `json:"data"`
}
