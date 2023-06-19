package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/api_key"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type GetAPIKeysResponse struct {
	common.BaseErrorResponse
	Data api_key.GetAPIKeysResponseWrapper `json:"data"`
}

type GetAPIKeyBaseInformationResponse struct {
	common.BaseErrorResponse
	Data api_key.GetAPIKeyBaseResponse `json:"data"`
}
