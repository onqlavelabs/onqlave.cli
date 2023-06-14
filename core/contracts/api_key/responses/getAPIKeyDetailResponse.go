package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/api_key"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type GetAPIKeyDetailResponse struct {
	common.BaseErrorResponse
	Data contracts.APIKeyDetail `json:"data"`
}

type GetAPIKeySensitiveInfoResponse struct {
	common.BaseErrorResponse
	Data contracts.APIKeySensitive `json:"data"`
}
